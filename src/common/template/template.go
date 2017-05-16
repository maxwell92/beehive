package template

import (
	"github.com/maxwell92/gokits/placeholder"
)

type Templator interface {
	Setup() interface{}
	Replace() error
}

type Template struct {
	Ph placeholder.PlaceHolder
}

func (t *Template) Setup() *Template {

}

func (t *Template) Replace() error {
	// t.Ph.Replace()
}
