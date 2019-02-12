package model

import (
	"config"
	"constant"
	"fmt"
	"strings"

	"github.com/imroc/req"
	"gopkg.in/mgo.v2/bson"
)

// 微信服务器推送过来的消息(事件)的通用消息头.
type MsgHeader struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string `xml:"MsgType"      json:"MsgType"`
}

// 微信服务器推送过来的消息(事件)的合集.
type EventMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MsgHeader
	EventType string `xml:"Event" json:"Event"`

	MsgId        int64   `xml:"MsgId"        json:"MsgId"`
	Content      string  `xml:"Content"      json:"Content"`
	MediaId      string  `xml:"MediaId"      json:"MediaId"`
	PicURL       string  `xml:"PicUrl"       json:"PicUrl"`
	Format       string  `xml:"Format"       json:"Format"`
	Recognition  string  `xml:"Recognition"  json:"Recognition"`
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"   json:"Location_X"`
	LocationY    float64 `xml:"Location_Y"   json:"Location_Y"`
	Scale        int     `xml:"Scale"        json:"Scale"`
	Label        string  `xml:"Label"        json:"Label"`
	Title        string  `xml:"Title"        json:"Title"`
	Description  string  `xml:"Description"  json:"Description"`
	URL          string  `xml:"Url"          json:"Url"`
	EventKey     string  `xml:"EventKey"     json:"EventKey"`
	Ticket       string  `xml:"Ticket"       json:"Ticket"`
	Latitude     float64 `xml:"Latitude"     json:"Latitude"`
	Longitude    float64 `xml:"Longitude"    json:"Longitude"`
	Precision    float64 `xml:"Precision"    json:"Precision"`

	// menu
	MenuId       int64 `xml:"MenuId" json:"MenuId"`
	ScanCodeInfo *struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`
		ScanResult string `xml:"ScanResult" json:"ScanResult"`
	} `xml:"ScanCodeInfo,omitempty" json:"ScanCodeInfo,omitempty"`
	SendPicsInfo *struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo,omitempty" json:"SendPicsInfo,omitempty"`
	SendLocationInfo *struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"`
		LocationY float64 `xml:"Location_Y" json:"Location_Y"`
		Scale     int     `xml:"Scale"      json:"Scale"`
		Label     string  `xml:"Label"      json:"Label"`
		PoiName   string  `xml:"Poiname"    json:"Poiname"`
	} `xml:"SendLocationInfo,omitempty" json:"SendLocationInfo,omitempty"`

	MsgID    int64  `xml:"MsgID"  json:"MsgID"`  // template, mass
	Status   string `xml:"Status" json:"Status"` // template, mass
	*mass           // mass
	*account        // account
	*dkf            // dkf
	*poi            // poi
	*card           // card
	*bizwifi        // bizwifi
	*file           // MsgType is file

	// shakearound
	ChosenBeacon *struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"ChosenBeacon,omitempty" json:"ChosenBeacon,omitempty"`
	AroundBeacons []struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"AroundBeacons>AroundBeacon,omitempty" json:"AroundBeacons,omitempty"`
}

type mass struct {
	//MsgID       int64  `xml:"MsgID"       json:"MsgID"`
	//Status      string `xml:"Status"      json:"Status"`
	TotalCount  int `xml:"TotalCount"  json:"TotalCount"`
	FilterCount int `xml:"FilterCount" json:"FilterCount"`
	SentCount   int `xml:"SentCount"   json:"SentCount"`
	ErrorCount  int `xml:"ErrorCount"  json:"ErrorCount"`
}

type account struct {
	ExpiredTime int64  `xml:"ExpiredTime" json:"ExpiredTime"`
	FailTime    int64  `xml:"FailTime"    json:"FailTime"`
	FailReason  string `xml:"FailReason"  json:"FailReason"`
}

type dkf struct {
	KfAccount     string `xml:"KfAccount"     json:"KfAccount"`
	FromKfAccount string `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string `xml:"ToKfAccount"   json:"ToKfAccount"`
}

type poi struct {
	UniqId string `xml:"UniqId" json:"UniqId"`
	PoiId  int64  `xml:"PoiId"  json:"PoiId"`
	Result string `xml:"Result" json:"Result"`
	Msg    string `xml:"Msg"    json:"Msg"`
}

type card struct {
	CardId              string `xml:"CardId"              json:"CardId"`
	RefuseReason        string `xml:"RefuseReason"        json:"RefuseReason"`
	IsGiveByFriend      int    `xml:"IsGiveByFriend"      json:"IsGiveByFriend"`
	FriendUserName      string `xml:"FriendUserName"      json:"FriendUserName"`
	UserCardCode        string `xml:"UserCardCode"        json:"UserCardCode"`
	OldUserCardCode     string `xml:"OldUserCardCode"     json:"OldUserCardCode"`
	ConsumeSource       string `xml:"ConsumeSource"       json:"ConsumeSource"`
	OuterId             int64  `xml:"OuterId"             json:"OuterId"`
	LocationName        string `xml:"LocationName"        json:"LocationName"`
	StaffOpenId         string `xml:"StaffOpenId"         json:"StaffOpenId"`
	VerifyCode          string `xml:"VerifyCode"          json:"VerifyCode"`
	RemarkAmount        string `xml:"RemarkAmount"        json:"RemarkAmount"`
	OuterStr            string `xml:"OuterStr"            json:"OuterStr"`
	Detail              string `xml:"Detail"              json:"Detail"`
	IsReturnBack        int    `xml:"IsReturnBack"        json:"IsReturnBack"`
	IsChatRoom          int    `xml:"IsChatRoom"          json:"IsChatRoom"`
	IsRestoreMemberCard int    `xml:"IsRestoreMemberCard" json:"IsRestoreMemberCard"`
	IsRecommendByFriend int    `xml:"IsRecommendByFriend" json:"IsRecommendByFriend"`
	PageId              string `xml:"PageId"              json:"PageId"`
	OrderId             string `xml:"OrderId"             json:"OrderId"`
}

type bizwifi struct {
	ConnectTime int64  `xml:"ConnectTime" json:"ConnectTime"`
	ExpireTime  int64  `xml:"ExpireTime"  json:"ExpireTime"`
	VendorId    string `xml:"VendorId"    json:"VendorId"`
	PlaceId     int64  `xml:"PlaceId"     json:"PlaceId"`
	DeviceNo    string `xml:"DeviceNo"    json:"DeviceNo"`
}

type file struct {
	FileKey      string `xml:"FileKey"      json:"FileKey"`
	FileMd5      string `xml:"FileMd5"      json:"FileMd5"`
	FileTotalLen string `xml:"FileTotalLen" json:"FileTotalLen"`
}

type EventFunc func(EventMsg) (interface{}, error)

var (
	urlJoinPhsMpGroup  string
	urlPhsGetGroupInfo string
)

func init() {
	if config.Conf.AppInfo.Env == "prod" {
		urlJoinPhsMpGroup = constant.URLPhsMPJoinGroup
		urlPhsGetGroupInfo = constant.URLPhsMPGetGroupInfo
	} else {
		urlJoinPhsMpGroup = constant.URLPhsMPJoinGroupDev
		urlPhsGetGroupInfo = constant.URLPhsMPGetGroupInfoDev
	}
}

func EventSubscribe(msg EventMsg) (interface{}, error) {
	user, _ := CreateUser(msg.FromUserName)
	if msg.EventKey != "" {
		scene, err := msg.Scene()
		if err != nil {
			return nil, err
		}
		var resData interface{}

		// 扫码关注场景，推送
		// TODO 服务注册形式
		switch {
		case strings.HasPrefix(scene, constant.WechatEventKeyScanCodeJoinPhsMPGroupPrefix):
			// 关注"写了吗" 小程序
			code := scene[len(constant.WechatEventKeyScanCodeJoinPhsMPGroupPrefix)+1:]
			go joinPhsMpGroup(user.Unionid, code)

			groupInfo, _ := getGroupInfoByCode(code)
			content := "你已成功加入班级"
			if nickname, ok := groupInfo["nickname"]; ok && nickname != "" {
				content = fmt.Sprintf("你已成功加入班级：%s", nickname)
			}
			resData = WechatTextRes{
				ToUserName:   msg.FromUserName,
				FromUserName: msg.ToUserName,
				CreateTime:   msg.CreateTime,
				Content:      content,
				MsgType:      "text",
			}
		}
		return resData, nil
	} else {
		return ResWelcomText(msg)
	}
}

func getGroupInfoByCode(code string) (map[string]string, error) {
	param := req.Param{
		"code": code,
	}
	r, err := req.Get(urlPhsGetGroupInfo, param)
	if err != nil {
		return nil, err
	}
	resData := struct {
		Status int               `json:"status"`
		Data   map[string]string `json:"data"`
	}{}
	err = r.ToJSON(&resData)
	return resData.Data, err
}

func EventUnSubscribe(msg EventMsg) (interface{}, error) {
	return nil, SetUserToUnFollow(msg.FromUserName)
}

func EventScan(msg EventMsg) (interface{}, error) {
	unionid, _ := GetUserUnionid(msg.FromUserName)
	if unionid == "" {
		return nil, nil
	}
	// SCAN 事件时 EventKey 即为 scene_str or scene_id
	scene := msg.EventKey
	var resData interface{}
	switch {
	case strings.HasPrefix(scene, constant.WechatEventKeyScanCodeJoinPhsMPGroupPrefix):
		// 关注"写了吗" 小程序
		code := scene[len(constant.WechatEventKeyScanCodeJoinPhsMPGroupPrefix)+1:]
		go joinPhsMpGroup(unionid, code)

		groupInfo, _ := getGroupInfoByCode(code)
		content := "你已成功加入班级"
		if nickname, ok := groupInfo["nickname"]; ok && nickname != "" {
			content = fmt.Sprintf("你已成功加入班级：%s", nickname)
		}
		resData = WechatTextRes{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   msg.CreateTime,
			Content:      content,
			MsgType:      "text",
		}
	}
	return resData, nil
}

func joinPhsMpGroup(unionid, groupCode string) error {
	query := bson.M{
		"unionid": unionid,
	}
	user, _ := findUser(query, defaultSelector)
	reqData := map[string]interface{}{
		"code":      groupCode,
		"unionId":   unionid,
		"nickName":  user.Nickname,
		"gender":    user.Sex,
		"city":      user.City,
		"province":  user.Province,
		"country":   user.Country,
		"avatarUrl": user.HeadImageURL,
		"language":  user.Language,
	}

	_, err := req.Post(urlJoinPhsMpGroup, req.BodyJSON(&reqData))
	return err
}

func EventLocation(msg EventMsg) (interface{}, error) {
	return nil, nil
}

func EventClick(msg EventMsg) (interface{}, error) {
	return nil, nil
}

func EventView(msg EventMsg) (interface{}, error) {
	return nil, nil
}

// Scene 获取二维码参数
func (this EventMsg) Scene() (scene string, err error) {
	const prefix = "qrscene_"
	if !strings.HasPrefix(this.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %s 为前缀: %s", prefix, this.EventKey)
		return
	}
	scene = this.EventKey[len(prefix):]
	return
}
