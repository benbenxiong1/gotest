package verifycode

import (
	"gotest/pkg/app"
	"gotest/pkg/config"
	"gotest/pkg/redis"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 设置redis验证码
func (s *RedisStore) Set(id string, value string) bool {
	// 过期时间
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return s.RedisClient.Set(s.KeyPrefix+id, value, ExpireTime)
}

// Get 获取验证码对应的值 可删除
func (s *RedisStore) Get(id string, clear bool) string {
	id = s.KeyPrefix + id
	val := s.RedisClient.Get(id)
	if clear {
		s.RedisClient.Del(id)
	}
	return val
}

// Verify 验证验证码是否正确
func (s *RedisStore) Verify(id string, answer string, clear bool) bool {
	val := s.Get(id, clear)
	return val == answer
}
