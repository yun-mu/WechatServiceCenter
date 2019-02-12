package view

import (
	"controller"

	"github.com/labstack/echo"
)

func initUserView(group *echo.Group) {
	group.GET("/status/follow", controller.GetUserFollowStatus)
	group.POST("/init", controller.InitUser)
}
