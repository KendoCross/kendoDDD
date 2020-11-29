package helper

import (
	"fmt"
	"net/smtp"
)

const (
	// 邮件服务器地址
	SMTP_MAIL_HOST = "smtp.163.com"
	// 端口
	SMTP_MAIL_PORT = "25"
	// 发送邮件用户账号
	SMTP_MAIL_USER = "cio_cheny@163.com"
	// 授权密码
	SMTP_MAIL_PWD = "SJAFMLTRRWAXXJIX" //cheny123
	// 发送邮件昵称
	SMTP_MAIL_NICKNAME = "王者之剑"
)

func SendMail(address []string, subject string, body string) (err error) {
	// 通常身份应该是空字符串，填充用户名.
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	contentType := "Content-Type: text/html; charset=UTF-8"
	for _, v := range address {
		s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
			v, SMTP_MAIL_NICKNAME, SMTP_MAIL_USER, subject, contentType, body)
		msg := []byte(s)
		addr := fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT)
		err = smtp.SendMail(addr, auth, SMTP_MAIL_USER, []string{v}, msg)
		if err != nil {
			return err
		}
	}
	return
}
