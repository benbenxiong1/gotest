package verifycode

import (
	"gotest/pkg/app"
	"gotest/pkg/config"
	"gotest/pkg/helpers"
	"gotest/pkg/logger"
	"gotest/pkg/redis"
	"gotest/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

// 单例
var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})

	return internalVerifyCode
}

func (v *VerifyCode) SendSMS(phone string) bool {
	code := v.generateVerifyCode(phone)
	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return sms.NewSMS().Send(phone, sms.Message{
		Data:     map[string]string{"code": code},
		Template: config.GetString("sms.aliyun.template_code"),
	})
}

// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (v *VerifyCode) CheckAnswer(key string, answer string) bool {

	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	// 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	if !app.IsProduction() &&
		(strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
			strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}

	return v.Store.Verify(key, answer, false)
}

func (v *VerifyCode) generateVerifyCode(key string) string {
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	// 为方便开发，本地环境使用固定验证码
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}
	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	v.Store.Set(key, code)
	return code
}
