package notifications

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

type Sender struct {
	SmtpHost string
	Email    string
	Password string
	Port     string
}

func (s Sender) Send(to []string, message string) error {
	auth := LoginAuth(s.Email, s.Password)
	err := smtp.SendMail(s.SmtpHost+":"+s.Port, auth, s.Email, to, []byte(message))
	return err
}

func (s Sender) SendHtml(senderName, subject, html string, to []string, cc []string, bcc []string) error {
	m := Mail{
		Sender:  s.Email,
		To:      to,
		Cc:      cc,
		Bcc:     bcc,
		Subject: subject,
		Body:    html,
	}
	receivers := []string{}
	receivers = append(receivers, to...)
	if len(cc) > 0 {
		receivers = append(receivers, cc...)
	}
	if len(bcc) > 0 {
		receivers = append(receivers, bcc...)
	}
	return s.Send(receivers, m.BuildHtml())
}

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

func (mail Mail) BuildHtml() string {

	msg := ""
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)

	if len(mail.To) > 0 {
		msg += fmt.Sprintf("To: %s\r\n", mail.To[0])
	}

	if len(mail.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	if len(mail.Bcc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Bcc, ";"))
	}

	msg += fmt.Sprintf("Content-Type: %s\r\n", "text/html; charset=UTF-8")
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
