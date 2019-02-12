package model

import (
	"constant"
	"fmt"
	"model/db"

	"github.com/imroc/req"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	// 状态:  0 表示没有关注公众号，5 表示已经关注
	Status int `bson:"status" json:"status"`

	IsSubscriber int    `bson:"subscribe" json:"subscribe"` // 用户是否订阅该公众号标识, 值为0时, 代表此用户没有关注该公众号, 拉取不到其余信息
	Openid       string `bson:"openid" json:"openid"`       // 用户的标识, 对当前公众号唯一
	Nickname     string `bson:"nickname" json:"nickname"`   // 用户的昵称
	Sex          int    `bson:"sex" json:"sex"`             // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	Language     string `bson:"language" json:"language"`   // 用户的语言, zh_CN, zh_TW, en
	City         string `bson:"city" json:"city"`           // 用户所在城市
	Province     string `bson:"province" json:"province"`   // 用户所在省份
	Country      string `bson:"country" json:"country"`     // 用户所在国家

	// 用户头像, 最后一个数值代表正方形头像大小(有0, 46, 64, 96, 132数值可选, 0代表640*640正方形头像), 用户没有头像时该项为空
	HeadImageURL string `bson:"headimgurl" json:"headimgurl"`

	SubscribeTime int64  `bson:"subscribe_time" json:"subscribe_time"`       // 用户关注时间, 为时间戳. 如果用户曾多次关注, 则取最后关注时间
	Unionid       string `bson:"unionid,omitempty" json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后, 才会出现该字段.
	Remark        string `bson:"remark" json:"remark"`                       // 公众号运营者对粉丝的备注, 公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupID       int64  `bson:"groupid" json:"groupid"`                     // 用户所在的分组ID

	TagIDList      []int  `bson:"tagid_list" json:"tagid_list"`           // Tag List
	SubscribeScene string `bson:"subscribe_scene" json:"subscribe_scene"` // 返回用户关注的渠道来源
	QrScene        int    `bson:"qr_scene" json:"qr_scene"`               // 二维码扫码场景（开发者自定义）场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
	QrSceneStr     string `bson:"qr_scene_str" json:"qr_scene_str"`       // 二维码扫码场景描述（开发者自定义）场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
}

// 获取用户列表返回的数据结构
type ListResult struct {
	TotalCount int `json:"total"` // 关注该公众账号的总用户数
	ItemCount  int `json:"count"` // 拉取的OPENID个数, 最大值为10000

	Data struct {
		OpenIdList []string `json:"openid,omitempty"`
	} `json:"data"` // 列表数据, OPENID的列表

	// 拉取列表的最后一个用户的OPENID, 如果 next_openid == "" 则表示没有了用户数据
	NextOpenId string `json:"next_openid"`
}

func InitUser() {
	var nextOpenID string
	var url string
	total := 0
	status := true
	for status {
		if nextOpenID == "" {
			url = "https://api.weixin.qq.com/cgi-bin/user/get?access_token="
		} else {
			url = "https://api.weixin.qq.com/cgi-bin/user/get?next_openid=" + nextOpenID + "&access_token="
		}
		url += accessToken

		result := ListResult{}
		r, err := req.Get(url)
		if err != nil {
			break
		}
		err = r.ToJSON(&result)
		if err != nil {
			break
		}

		for _, openid := range result.Data.OpenIdList {
			CreateUser(openid)
		}

		total += len(result.Data.OpenIdList)
		if total >= result.TotalCount {
			status = false
		}
		nextOpenID = result.NextOpenId
	}
}

func CreateUser(openid string) (user User, err error) {
	query := bson.M{
		"openid": openid,
	}

	selector := bson.M{
		"status":     1,
		"nickname":   1,
		"headimgurl": 1,
	}

	oldUser, err := findUser(query, selector)
	user, _ = GetWeixinUserInfo(openid)
	if err != nil {
		// 没有关注，且数据库没有存储信息
		if user.IsSubscriber != 0 {
			user.ID = bson.NewObjectId()
			user.Status = constant.UserFollowStatus
			insertUsers(user)
		}
		err = nil
		return
	}
	var update bson.M

	if oldUser.Status == constant.UserFollowStatus {
		if oldUser.Nickname != user.Nickname || oldUser.HeadImageURL != user.HeadImageURL {
			update = bson.M{
				"$set": bson.M{
					"nickname":   user.Nickname,
					"headimgurl": user.HeadImageURL,
					"sex":        user.Sex,
					"language":   user.Language,
					"country":    user.Country,
					"city":       user.City,
					"province":   user.Province,
				},
			}
		}
	} else {
		update = bson.M{
			"$set": bson.M{
				"status":     constant.UserFollowStatus,
				"nickname":   user.Nickname,
				"headimgurl": user.HeadImageURL,
				"sex":        user.Sex,
				"language":   user.Language,
				"country":    user.Country,
				"city":       user.City,
				"province":   user.Province,
			},
		}
	}
	err = updateUser(query, update)
	return
}

func SetUserToUnFollow(openid string) error {
	query := bson.M{
		"openid": openid,
	}

	update := bson.M{
		"$set": bson.M{
			"status": constant.UserUnFollowStatus,
		},
	}
	return updateUser(query, update)
}

func GetUserOpenid(unionid string) (string, error) {
	query := bson.M{
		"unionid": unionid,
		"status": bson.M{
			"$gte": constant.UserFollowStatus,
		},
	}
	selector := bson.M{
		"openid": 1,
	}
	user, err := findUser(query, selector)
	return user.Openid, err
}

func GetUserUnionid(openid string) (string, error) {
	query := bson.M{
		"openid": openid,
	}
	selector := bson.M{
		"unionid": 1,
	}
	user, err := findUser(query, selector)
	return user.Unionid, err
}

func GetWeixinUserInfo(openid string) (user User, err error) {
	url := fmt.Sprintf(constant.URLWechatUserInfo, accessToken, openid)
	r, _ := req.Get(url)
	err = r.ToJSON(&user)
	return
}

/****************************************** user basic action ****************************************/

func findUser(query, selector interface{}) (User, error) {
	data := User{}
	cntrl := db.NewCloneMgoDBCntlr()
	defer cntrl.Close()
	table := cntrl.GetTable(constant.TableUser)
	err := table.Find(query).Select(selector).One(&data)
	return data, err
}

func updateUser(query, update interface{}) error {
	return updateDoc(constant.TableUser, query, update)
}

func insertUsers(docs ...interface{}) error {
	return insertDocs(constant.TableUser, docs...)
}

/****************************************** user redis action ****************************************/
