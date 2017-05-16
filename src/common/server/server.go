package server

import (
	"github.com/kataras/iris"
)

type Server interface {
	Up(string)
}

func Instance() Server {
	is := new(IrisServer)
	return is
}

type IrisServer struct {
}

func (i IrisServer) Up(host string) {
	iris.Listen(host)
}



