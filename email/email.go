package email

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/topicai/candy"
)

// adim email info and SMTP address
type SmtpInfo struct {
	AdminEmail  string //email address
	AdminSecrt  string //email secrt
	ESMTPServer string // SMTP host address, format: host:port
	subject     string //subject of email
	text        string //text of email
	ch          chan *gomail.Message
}

func NewSmtpInfo(e, a, s string) *SmtpInfo {
	return &SmtpInfo{
		ESMTPServer: e,
		AdminEmail:  a,
		AdminSecrt:  s,
		subject:     "k8s key and crt",
		text:        "Successful: your crt and key for k8s are in attachment!\n",
		ch:          make(chan *gomail.Message),
	}
}

func (smtp *SmtpInfo) SMTPSvcPool() {
	// SMTP host and port
	host, port, _ := net.SplitHostPort(smtp.ESMTPServer)
	portint, err := strconv.Atoi(port)
	fmt.Println("smtp host:", host)
	fmt.Println("smtp port:", portint)
	fmt.Println("smtp info:", smtp)

	candy.Must(err)

	d := gomail.NewDialer(host, portint, smtp.AdminEmail, smtp.AdminSecrt)

	var s gomail.SendCloser
	open := false

	for {
		select {
		case m, ok := <-smtp.ch:
			if !ok {
				// channel closed
				return
			}

			if !open {
				// dial to  SMTP
				s, err = d.Dial()
				fmt.Println("ssl_flag=%v open_flag=%v err=%v", d.SSL, open, err)
				candy.Must(err)
				open = true
			}

			//send email
			err := gomail.Send(s, m)
			candy.Must(err)

			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
		case <-time.After(30 * time.Second):
			if open {
				err := s.Close()
				fmt.Println("send email timeout")
				candy.Must(err)
				open = false
			}
		}
	}
}

func (s *SmtpInfo) SendEmail(to, crt, key string) {
	m := gomail.NewMessage()
	m.SetHeader("From", s.AdminEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", s.subject)
	m.SetBody("text/plan", s.text)
	m.Attach(crt)
	m.Attach(key)

	s.ch <- m

	fmt.Println("Send email...")
	return
}
