package untils

import (
	"github.com/zhanghanchen1014/viper/pkg"
	"log"
)

type SendSms interface {
	ALiYunSms(mobile, code string) error
}

type ALiYun struct {
}

func (s *ALiYun) ALiYunSms(mobile, code string) error {
	sms, err := pkg.SendSms(mobile, code)
	if err != nil {
		return err
	}
	log.Println(sms)
	return err
}
