package routes

import (
	"github.com/gin-gonic/gin"
	"gotest/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterAPIRoutes(e *gin.Engine) {
	v1 := e.Group("v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 发送短信
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
		}

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
