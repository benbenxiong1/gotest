package verifycode

type Store interface {
	// Set 设置验证码
	Set(id string, value string) bool
	// Get 查询验证码
	Get(id string, clear bool) string
	// Verify 检查验证码
	Verify(id, answer string, clear bool) bool
}
