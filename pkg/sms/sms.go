package sms

import (
	"gotest/pkg/config"
	"sync"
)

type Message struct {
	Data     map[string]string
	Template string
	Content  string
}

type SMS struct {
	Driver Driver
}

// once 单例
var once sync.Once

// internalSMS 内部使用的SMS变量
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}

// Send 发送
func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
