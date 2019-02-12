package controller

import (
	"model"

	"github.com/labstack/echo"
)

/**
 * @apiDefine GetUserFollowStatus GetUserFollowStatus
 * @apiDescription 获取是否关注了服务号
 *
 * @apiParam {String} unionid unionid
 *
 * @apiParamExample  {json} Request-Example:
 *      {
 *        "unionid":"Unionid",
 *      }
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *        "is_follow": Boolean, // 是否关注了服务号
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
 * @api {get} /api/v1/user/status/follow GetUserFollowStatus
 * @apiVersion 1.0.0
 * @apiName GetUserFollowStatus
 * @apiGroup User
 * @apiUse GetUserFollowStatus
 */
func GetUserFollowStatus(c echo.Context) error {
	unionid := c.QueryParam("unionid")
	openid, _ := model.GetUserOpenid(unionid)
	resData := map[string]bool{
		"is_follow": openid != "",
	}
	return retRawJSONData(c, resData)
}

func InitUser(c echo.Context) error {
	model.InitUser()
	return retRawJSONData(c, "ok")
}

func writeUserLog(funcName, errMsg string, err error) {
	writeLog("user.go", funcName, errMsg, err)
}
