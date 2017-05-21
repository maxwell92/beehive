package node

import (
	"controller"
)

type NodeController struct {
	controller.Controller
}

func (nc NodeController) Get() {
	nc.WriteOk("node get ok")
}

func (nc NodeController) Post() {
	nc.WriteOk("node post ok")
}

