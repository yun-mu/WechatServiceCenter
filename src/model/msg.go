package model

import (
	"config"
	"constant"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"sort"

	"github.com/imroc/req"
)

type Template struct {
	ToUser      string                 `json:"touser"`                // 必须, 接受者OpenID
	TemplateID  string                 `json:"template_id"`           // 必须, 模版ID
	URL         string                 `json:"url,omitempty"`         // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	MiniProgram *MiniProgram           `json:"miniprogram,omitempty"` // 可选, 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        map[string]interface{} `json:"data"`                  // 必须, 模板数据
}

type MiniProgram struct {
	AppID    string `json:"appid"`    // 必选; 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
	PagePath string `json:"pagepath"` // 必选; 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
}

type TemplateResult struct {
	Error
	MsgID int64 `json:"msgid"`
}

type Error struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WechatTextRes struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string   `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string   `xml:"MsgType"      json:"MsgType"`
	Content      string   `xml:"Content" json:"Content"` // 回复的消息内容(换行: 在content中能够换行, 微信客户端支持换行显示)
}

func SendTemplate(template Template) (TemplateResult, error) {
	url := fmt.Sprintf(constant.URLWechatTemplateSend, accessToken)
	resData := TemplateResult{}
	r, err := req.Post(url, req.BodyJSON(&template))
	if err != nil {
		return resData, err
	}
	err = r.ToJSON(&resData)
	return resData, err
}

func GetSignature(timestamp, nonce string) string {
	return GetSign(config.Conf.Wechat.Token, timestamp, nonce)
}

// Sign 微信公众号 url 签名.
func GetSign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

func ResWelcomText(msg EventMsg) (interface{}, error) {
	return WechatTextRes{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		Content:      "终于等到你！",
		MsgType:      "text",
	}, nil
}

func ResCommonText(msg EventMsg) (interface{}, error) {
	return WechatTextRes{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		Content:      "Sorry，此功能暂未开发。",
		MsgType:      "text",
	}, nil
}
