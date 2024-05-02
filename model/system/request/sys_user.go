package request

import "gvaTemplate/model/system"

type ChangePassword struct {
	UserId      uint
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type CaptchaRegister struct {
	system.User
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
