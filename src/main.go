/*
Package main package is the entry file
*/
package main

import (
	"config"
	"constant"
	"controller"

	mid "middleware"
	"view"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/robfig/cron"
	validator "gopkg.in/go-playground/validator.v9"
)

func main() {
	go startTimer()
	startWeb()
}

func startWeb() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Validator = &mid.DefaultValidator{Validator: validator.New()}

	v1 := e.Group(constant.APIPrefix)
	view.InitViewV1(v1)

	e.Logger.Fatal(e.Start(config.Conf.AppInfo.Addr))
}

func startTimer() {
	c := cron.New()
	controller.StartHourTimer()
	c.AddFunc(constant.TimerEveryHour, controller.StartHourTimer)
	c.Start()
}
