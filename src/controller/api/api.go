package api

import (
	 "controller"
)

type Api struct {
	controller.Controller
}

func (a Api) Get() {
	a.Write("Hello World!")
}
