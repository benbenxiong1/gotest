package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gotest/app/http/controllers/api/v1"
	"gotest/app/models/user"
	"gotest/app/requests"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

func (s *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// 检查数据 并返回数据
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (s SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
