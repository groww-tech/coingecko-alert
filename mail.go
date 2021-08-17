package main

import (
	"gopkg.in/gomail.v2"
)

type mail struct {
	sender gomail.SendCloser
}

func newMail(host string, port int, username, password string) *mail {
	d := gomail.NewDialer(host, port, username, password)
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	return &mail{s}
}

func (t *mail) notify(from string, to string, message string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Coingecko Notify!")
	m.SetBody("text/html", message)

	if err := gomail.Send(t.sender, m); err != nil {
		panic(err)
	}
}
