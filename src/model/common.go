package model

import (
	"config"
	"constant"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/imroc/req"
)

type weixinAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

type weixinJSApiTicketResponse struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
}

type QrcodeParam struct {
	ExpireSeconds int    `json:"expire_seconds"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneStr string `json:"scene_str"`
			SceneId  int32  `json:"scene_id"`
		} `json:"scene"`
	} `json:"action_info"`
}

type QrcodeRes struct {
	ExpireSeconds int    `json:"expire_seconds,omitempty"`
	Ticket        string `json:"ticket"`
	URL           string `json:"url"`
}

var (
	appid     string
	appSecret string

	accessToken string
	jsapiTicket string
	letterRunes []rune

	jsapiTicketMutex sync.RWMutex
	tokenMutex       sync.RWMutex
)

func init() {
	rand.Seed(time.Now().UnixNano())

	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	appid = config.Conf.Wechat.AppID
	appSecret = config.Conf.Wechat.AppSecret
}

func CreateQrcode(data QrcodeParam) (QrcodeRes, error) {
	url := fmt.Sprintf(constant.URLWechatCreateQrcode, accessToken)
	r, _ := req.Post(url, req.BodyJSON(&data))
	res := QrcodeRes{}
	err := r.ToJSON(&res)
	return res, err
}

func GetAccessToken() string {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	return accessToken
}

func GetJSApiTicket() string {
	jsapiTicketMutex.RLock()
	defer jsapiTicketMutex.RUnlock()
	return jsapiTicket
}

func GetJSSDKConfig(url string) map[string]string {
	nonceStr := randStringRunes(32)
	timestamp := time.Now().Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)

	jsapiTicketMutex.RLock()
	defer jsapiTicketMutex.RUnlock()
	sign := wxConfigSign(jsapiTicket, nonceStr, timestampStr, url)

	return map[string]string{
		"appId":     appid,
		"nonce_str": nonceStr,
		"signature": sign,
		"timestamp": timestampStr,
	}
}

func UpdateAccessToken() error {
	resp, err := requestAccessToken()
	if err != nil {
		return err
	}

	if resp.AccessToken == "" {
		return fmt.Errorf("errcode: %d, errmsg: %s", resp.Errcode, resp.Errmsg)
	}
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	accessToken = resp.AccessToken
	return nil
}

func UpdateJSApiTicket() error {
	if accessToken == "" {
		return errors.New("accessToken 为空")
	}
	resp, err := requestJSApiTicket(accessToken)
	if err != nil {
		return err
	}
	jsapiTicketMutex.Lock()
	defer jsapiTicketMutex.Unlock()
	jsapiTicket = resp.Ticket
	return nil
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func wxConfigSign(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	if i := strings.IndexByte(url, '#'); i >= 0 {
		url = url[:i]
	}

	n := len("jsapi_ticket=") + len(jsapiTicket) +
		len("&noncestr=") + len(nonceStr) +
		len("&timestamp=") + len(timestamp) +
		len("&url=") + len(url)
	buf := make([]byte, 0, n)

	buf = append(buf, "jsapi_ticket="...)
	buf = append(buf, jsapiTicket...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, nonceStr...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, timestamp...)
	buf = append(buf, "&url="...)
	buf = append(buf, url...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

func requestAccessToken() (weixinAccessTokenResponse, error) {
	url := fmt.Sprintf(constant.URLWechatAccessToken, appid, appSecret)
	r, err := req.Get(url)
	var respBody weixinAccessTokenResponse
	if err != nil {
		return respBody, err
	}
	err = r.ToJSON(&respBody)
	return respBody, err
}

func requestJSApiTicket(accessToken string) (weixinJSApiTicketResponse, error) {
	url := fmt.Sprintf(constant.URLWechatJSApiTicket, accessToken)
	r, err := req.Get(url)
	var respBody weixinJSApiTicketResponse
	if err != nil {
		return respBody, err
	}
	err = r.ToJSON(&respBody)
	return respBody, err
}

/****************************************** base redis action ****************************************/
