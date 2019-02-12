package controller

import (
	"net/http"
	"util/log"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ErrorRes ErrorResponse
type ErrorRes struct {
	Status int    `json:"status"`
	ErrMsg string `json:"err_msg"`
}

// DataRes DataResponse
type DataRes struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// RetError response error, wrong response
func retError(c echo.Context, code, status int, errMsg string) error {
	return c.JSON(code, ErrorRes{
		Status: status,
		ErrMsg: errMsg,
	})
}

// RetData response data, correct response
func retData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, DataRes{
		Status: 200,
		Data:   data,
	})
}

func retRawJSONData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func retXMLData(c echo.Context, data interface{}) error {
	return c.XML(http.StatusOK, data)
}

var logger = log.GetLogger()

func writeLog(fileName, funcName, errMsg string, err error) {
	logger.WithFields(logrus.Fields{
		"package":  "controller",
		"file":     fileName,
		"function": funcName,
		"err":      err,
	}).Warn(errMsg)
}
