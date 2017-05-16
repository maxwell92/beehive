package mail

import (

)

type Mailer interface {
	Setup() interface{}
	Send() error
}

type Mail struct {

}

func (m *Mail) Setup() *Mail {

}

func (m *Mail) Send() error {
}
