package email

import (
	"bytes"
	"fmt"
	"net"
	"net/smtp"
	"strconv"
	"strings"
)

type SmtpEmailClient struct {
	Username string
	Password string
	Host     string
	Port     int
}

func (c SmtpEmailClient) Send(recipients []string, subject string, body string) error {
	auth := smtp.PlainAuth("", c.Username, c.Password, c.Host)
	addr := net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("From: %s\r\n", c.Username))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(recipients, ",")))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	msg.WriteString(fmt.Sprintf("\r\n%s\r\n", body))
	return smtp.SendMail(addr, auth, c.Username, recipients, msg.Bytes())
}
