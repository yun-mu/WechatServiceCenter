package controller

import (
	"constant"
	"controller/param"
	"model"
	"net/http"

	"github.com/labstack/echo"
)

/**
 * @apiDefine GetAccessToken GetAccessToken
 * @apiDescription 获取冰岩在线服务号的access_token, 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/AccessToken
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "access_token": String,
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
 * @api {get} /api/v1/access-token GetAccessToken
 * @apiVersion 1.0.0
 * @apiName GetAccessToken
 * @apiGroup Index
 * @apiUse GetAccessToken
 */
func GetAccessToken(c echo.Context) error {
	resData := map[string]interface{}{
		"access_token": model.GetAccessToken(),
	}
	return retRawJSONData(c, resData)
}

/**
 * @apiDefine GetJSApiTicket GetJSApiTicket
 * @apiDescription 获取冰岩在线服务号的 jsapi_ticket, 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/JsApiTicket
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "jsapi_ticket": String,
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
 * @api {get} /api/v1/js-api-ticket GetJSApiTicket
 * @apiVersion 1.0.0
 * @apiName GetJSApiTicket
 * @apiGroup Index
 * @apiUse GetJSApiTicket
 */
func GetJSApiTicket(c echo.Context) error {
	resData := map[string]string{
		"jsapi_ticket": model.GetJSApiTicket(),
	}
	return retRawJSONData(c, resData)
}

/**
 * @apiDefine GetJSSDKConfig GetJSSDKConfig
 * @apiDescription 获取冰岩在线服务号的 JSSDKConfig, 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/signature
 *
 * @apiParam {String} url	分享的网站链接
 *
 * @apiParamExample  {query} Request-Example:
 *      {
 *        "url": String
 *      }
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "appId": String,
 *       "nonce_str": String,
 *       "signature": String,
 *       "timestamp": String,
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
 * @api {get} /api/v1/signature GetJSSDKConfig
 * @apiVersion 1.0.0
 * @apiName GetJSSDKConfig
 * @apiGroup Index
 * @apiUse GetJSSDKConfig
 */
func GetJSSDKConfig(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return retRawJSONData(c, "")
	}
	resData := model.GetJSSDKConfig(url)
	return retRawJSONData(c, resData)
}

/**
 * @apiDefine CreateJSSDKConfig CreateJSSDKConfig
 * @apiDescription 获取冰岩在线服务号的 JSSDKConfig (POST 版本), 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/signature
 *
 * @apiParam {String} url	分享的网站链接
 *
 * @apiParamExample  {json} Request-Example:
 *      {
 *        "url": String
 *      }
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "appId": String,
 *       "nonce_str": String,
 *       "signature": String,
 *       "timestamp": String,
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
 * @api {post} /api/v1/signature CreateJSSDKConfig
 * @apiVersion 1.0.0
 * @apiName CreateJSSDKConfig
 * @apiGroup Index
 * @apiUse CreateJSSDKConfig
 */
func CreateJSSDKConfig(c echo.Context) error {
	data := param.URLParam{}
	err := c.Bind(&data)
	if err != nil {
		writeIndexLog("CreateJSSDKConfig", constant.ErrorMsgParamWrong, err)
		return c.XML(http.StatusBadGateway, nil)
	}
	url := data.URL
	if url == "" {
		return retRawJSONData(c, "")
	}
	resData := model.GetJSSDKConfig(url)
	return retRawJSONData(c, resData)
}

/**
 * @apiDefine CreateQrcode CreateQrcode
 * @apiDescription 获取冰岩在线服务号的 二维码，具体请看微信公众号文档
 *
 * @apiParam {Number} expire_seconds
 * @apiParam {String} action_name
 * @apiParam {Object} action_info
 * @apiParam {Number} action_info.scene_id
 * @apiParam {String} action_info.scene_str
 *
 * @apiParamExample  {json} Request-Example:
 *      {
 *        "expire_seconds": Number,
 *        "action_name": String,
 *        "action_info": {
 *            "scene": {
 *	              "scene_str": "test",
 *                "scene_id": 123,
 *              },
 *          },
 *      }
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "ticket": String, // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
 *       "expire_seconds": Number, // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。
 *       "url": String, // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
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
 * @api {post} /api/v1/qrcode CreateQrcode
 * @apiVersion 1.0.0
 * @apiName CreateQrcode
 * @apiGroup Index
 * @apiUse CreateQrcode
 */
func CreateQrcode(c echo.Context) error {
	data := model.QrcodeParam{}
	err := c.Bind(&data)
	if err != nil {
		writeMsgLog("CreateQrcode", constant.ErrorMsgParamWrong, err)
		return retError(c, http.StatusBadRequest, http.StatusBadRequest, constant.ErrorMsgParamWrong)
	}
	resData, err := model.CreateQrcode(data)
	if err != nil {
		writeMsgLog("CreateQrcode", "创建二维码失败", err)
		return retError(c, http.StatusBadGateway, http.StatusBadGateway, "创建二维码失败")
	}
	return retRawJSONData(c, resData)
}

func writeIndexLog(funcName, errMsg string, err error) {
	writeLog("index.go", funcName, errMsg, err)
}
