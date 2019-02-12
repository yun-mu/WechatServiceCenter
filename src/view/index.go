/*
The index view contains the index view and other view that uses `/` to begin
*/
package view

import (
	"controller"

	"github.com/labstack/echo"
)

func InitViewV1(group *echo.Group) {
	group.GET("/access-token", controller.GetAccessToken)
	group.GET("/AccessToken", controller.GetAccessToken)
	group.GET("/js-api-ticket", controller.GetJSApiTicket)
	group.GET("/JsApiTicket", controller.GetJSApiTicket)
	group.GET("/signature", controller.GetJSSDKConfig)
	group.POST("/signature", controller.CreateJSSDKConfig)

	group.POST("/qrcode", controller.CreateQrcode)

	msg := group.Group("/msg")
	initMsgView(msg)

	user := group.Group("/user")
	initUserView(user)
}
