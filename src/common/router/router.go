package router

import (
	"github.com/kataras/iris"
	"controller/api"
)

type Router struct {
	Api *api.Api
}


func Setup() {
	r := &Router{
		Api: new(api.Api),
	}

	iris.API("/", *r.Api)
}
