package mail

import (
	"gopkg.in/gomail.v2"
)

const (
	// 邮件服务器地址
	MailHost = "smtp.qq.com"
	// 端口
	MailPort = 587
	// 发送邮件用户账号
	MailUser = "809979396@qq.com"
	// 授权密码
	MailPwd = "fffiowhhjbnkbfbc"
)

/*
title 使用gomail发送邮件
@param string body 邮件内容
@return error
*/
func SendGoMail(body string) error {
	var mailAddress []string
	mailAddress = append(mailAddress, "xiejiazheng@valleysound.xyz")

	m := gomail.NewMessage()
	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	nickname := "gomail"
	m.SetHeader("From", nickname+"<"+MailUser+">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", "bitcoin-node")
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(MailHost, MailPort, MailUser, MailPwd)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
