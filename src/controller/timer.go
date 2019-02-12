package controller

import "model"

func StartHourTimer() {
	// 先更新 access_token 后更新ticket
	err := model.UpdateAccessToken()
	if err != nil {
		writeTimerLog("UpdateAccessToken", "更新 access_token 失败", err)
	}
	err = model.UpdateJSApiTicket()
	if err != nil {
		writeTimerLog("UpdateJSApiTicket", "更新 js api ticket 失败", err)
	}
}

func writeTimerLog(funcName, errMsg string, err error) {
	writeLog("timer.go", funcName, errMsg, err)
}
