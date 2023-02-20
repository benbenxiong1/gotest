package auth

import (
	"fmt"
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
	type phoneExistRequest struct {
		phone string `json:"phone"`
	}

	request := phoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败 返回422状态码
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())

		// 请求错误 中断请求
		return
	}

	// 验证器验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	// 检查数据 并返回数据
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.phone),
	})
}
