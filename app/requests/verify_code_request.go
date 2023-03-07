package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gotest/pkg/captcha"
)

type VerifyCodePhoneRequest struct {
	CaptchaId     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Phone         string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone 验证表单，返回长度等于零即通过
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"phone":          []string{"required", "digits:11"},
	}

	// 2. 定制错误消息
	message := govalidator.MapData{
		"captcha_id": []string{
			"required:验证码的ID为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码的答案必填",
			"digits:图片验证码的长度为6",
		},
		"phone": []string{
			"required:手机号必传",
			"digits:手机号的长度为为11位",
		},
	}
	errs := validate(data, rules, message)

	_data := data.(*VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaId, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}
