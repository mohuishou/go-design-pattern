package template

import "fmt"

// ISMS ISMS
type ISMS interface {
	send(content string, phone int) error
}

// SMS 短信发送基类
type sms struct {
	ISMS
}

// Valid 校验短信字数
func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

// Send 发送短信
func (s *sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return err
	}

	// 调用子类的方法发送短信
	return s.send(content, phone)
}

// TelecomSms 走电信通道
type TelecomSms struct {
	*sms
}

// NewTelecomSms NewTelecomSms
func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	// 这里有点绕，是因为 go 没有继承，用嵌套结构体的方法进行模拟
	// 这里将子类作为接口嵌入父类，就可以让父类的模板方法 Send 调用到子类的函数
	// 实际使用中，我们并不会这么写，都是采用组合+接口的方式完成类似的功能
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (tel *TelecomSms) send(content string, phone int) error {
	fmt.Println("send by telecom success")
	return nil
}
