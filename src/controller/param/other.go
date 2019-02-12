package param

import "model"

/*---------------------------------- 其他 ---------------------------------------*/

type SignatureCheck struct {
	Signature string `json:"signature" query:"signature" form:"signature" validate:"required"`
	Timestamp string `json:"timestamp" query:"timestamp" form:"timestamp" validate:"required"`
	Nonce     string `json:"nonce" query:"nonce" form:"nonce" validate:"required"`
	Echostr   string `json:"echostr" query:"echostr" form:"echostr" validate:"required"`
}

type TemplatesParam struct {
	Templates []model.Template `json:"templates" query:"templates"`
}
