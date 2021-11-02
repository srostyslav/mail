package mail

import (
	gomail "gopkg.in/gomail.v2"
)

type mailSender struct {
	Login, Psw, Domain string
	Bcc                []string
	Port               int
}

var Mail = &mailSender{}

func (m *mailSender) Init(login, psw, domain string, port int, bcc []string) {
	m.Login, m.Psw, m.Bcc = login, psw, bcc
	m.Domain, m.Port = domain, port
}

func (a *mailSender) SendEmail(subject, text, fileName string, to []string) error {

	var err error

	for _, receipient := range append(to, a.Bcc...) {

		m := gomail.NewMessage()
		m.SetHeader("From", a.Login)
		m.SetHeader("To", receipient)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", text)

		if fileName != "" {
			m.Attach(fileName)
		}

		d := gomail.NewPlainDialer(a.Domain, a.Port, a.Login, a.Psw)
		if e := d.DialAndSend(m); e != nil {
			err = e
		}
	}

	return err
}
