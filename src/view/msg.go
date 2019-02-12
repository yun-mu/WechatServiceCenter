package view

import (
	"controller"

	"github.com/labstack/echo"
)

func initMsgView(group *echo.Group) {
	group.POST("/template/list/action/send", controller.SendTemplate)
	group.GET("/event", controller.DelSignature)
	group.POST("/event", controller.DelEvent)
}
