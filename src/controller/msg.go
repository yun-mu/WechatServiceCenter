package controller

import (
	"constant"
	"controller/param"
	"model"
	"net/http"

	"github.com/labstack/echo"
)

/**
 * @apiDefine SendTemplate SendTemplate
 * @apiDescription 发送模板消息
 *
 * @apiParam {[]Object} templates	模板数组
 * @apiParam {String} touser	接收者unionid
 * @apiParam {String} template_id	模板ID
 * @apiParam {String} url	模板跳转链接（海外帐号没有跳转能力）
 * @apiParam {String} miniprogram	跳小程序所需数据，不需跳小程序可不用传该数据
 * @apiParam {String} data	模板数据
 *
 * @apiParamExample  {json} Request-Example:
 *      {
 *        "templates": [{
 *            "touser":"Unionid",
 *            "template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
 *            "url":"http://weixin.qq.com/download",
 *            "miniprogram":{
 *                "appid":"xiaochengxuappid12345", // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
 *                "pagepath":"index?foo=bar" // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），暂不支持小游戏
 *              },
 *            "data":{
 *                   "keyword1":{
 *                       "value":"巧克力",
 *                       "color":"#173177" // 模板内容字体颜色，不填默认为黑色
 *                   },
 *                   "keyword2": {
 *                       "value":"39.8元",
 *                       "color":"#173177"
 *                   },
 *                   "keyword3": {
 *                       "value":"2014年9月22日",
 *                       "color":"#173177"
 *                   },
 *            }
 *        }]
 *      }
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 200,
 *       "data": [{
 *           "errcode": Number,
 *           "errmsg": String,
 *           "msgid": Number,
 *         }]
 *     }
 * @apiError {Number} status 状态码
 * @apiError {String} err_msg 错误信息
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 401 Unauthorized
 *     {
 *       "status": 401,
 *       "err_msg": "Unauthorized"
 *     }
 */
/**
 * @api {post} /api/v1/msg/template/list/action/send SendTemplate
 * @apiVersion 1.0.0
 * @apiName SendTemplate
 * @apiGroup Message
 * @apiUse SendTemplate
 */
func SendTemplate(c echo.Context) error {
	data := param.TemplatesParam{}
	err := c.Bind(&data)
	if err != nil {
		writeMsgLog("SendTemplate", constant.ErrorMsgParamWrong, err)
		return retError(c, http.StatusBadRequest, http.StatusBadRequest, constant.ErrorMsgParamWrong)
	}

	resData := make([]model.TemplateResult, len(data.Templates))
	for i, template := range data.Templates {
		openid, err := model.GetUserOpenid(template.ToUser)
		if openid == "" || err != nil {
			writeMsgLog("SendTemplate", "获取用户openid失败", err)
			continue
		}
		template.ToUser = openid
		resData[i], err = model.SendTemplate(template)
		if err != nil {
			writeMsgLog("SendTemplate", "发送模板失败", err)
			continue
		}
	}

	return retData(c, resData)
}

// DelEvent 处理微信的消息事件
func DelEvent(c echo.Context) error {
	data := model.EventMsg{}
	err := c.Bind(&data)
	if err != nil {
		writeMsgLog("DelEvent", constant.ErrorMsgParamWrong, err)
		return c.XML(http.StatusBadGateway, nil)
	}
	var resData interface{}
	if data.MsgType == "event" {
		// 处理事件消息
		event := map[string]model.EventFunc{
			constant.EventTypeSubscribe:   model.EventSubscribe,
			constant.EventTypeUnsubscribe: model.EventUnSubscribe,
			constant.EventTypeLocation:    model.EventLocation,
			constant.EventTypeClick:       model.EventClick,
			constant.EventTypeScan:        model.EventScan,
			constant.EventTypeVIEW:        model.EventView,
		}

		resData, err = event[data.EventType](data)
		if err != nil {
			writeMsgLog("DelEvent", "服务器错误", err)
			return c.XML(http.StatusBadGateway, nil)
		}
	} else {
		// 接收普通消息
		resData, _ = model.ResCommonText(data)
	}

	return c.XML(http.StatusOK, &resData)
}

// DelSignature 处理微信服务器的回调URL确认
func DelSignature(c echo.Context) error {
	data := param.SignatureCheck{}
	err := c.Bind(&data)
	echostr := ""
	if err != nil {
		writeMsgLog("DelSignature", constant.ErrorMsgParamWrong, err)
		return c.String(http.StatusOK, echostr)
	}
	signature := model.GetSignature(data.Timestamp, data.Nonce)
	if signature == data.Signature {
		echostr = data.Echostr
	}
	// oh, fork, 不能返回json
	return c.String(http.StatusOK, echostr)
}

func writeMsgLog(funcName, errMsg string, err error) {
	writeLog("msg.go", funcName, errMsg, err)
}
