package constant

const (

	/****************************************** wechat ****************************************/

	URLWechatAccessToken  = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	URLWechatJSApiTicket  = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	URLWechatTemplateSend = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
	URLWechatUserInfo     = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
	URLWechatCreateQrcode = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"

	URLPhsMPJoinGroup       = "<phs prefix>/api/unopen/group/action/join"
	URLPhsMPJoinGroupDev    = "<phs prefix>/api/unopen/group/action/join"
	URLPhsMPGetGroupInfo    = "<phs prefix>/api/unopen/group"
	URLPhsMPGetGroupInfoDev = "<phs prefix>/api/unopen/group"
)
