package api

import (
	"github.com/kataras/iris"
)

type Api struct {
	*iris.Context
}

func (a Api) Get() {
	a.Write("Hello World!")
}
