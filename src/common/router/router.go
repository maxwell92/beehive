package router

import (
	"controller/api"
	"controller/history"
	"controller/node"
	"controller/template"
	"github.com/kataras/iris"
)

type Router struct {
	Api *api.Api
	History *history.HistoryController
	Node *node.NodeController
	template *template.TemplateController
}

func Setup() {
	r := &Router{
		Api: new(api.Api),
		History: new(history.HistoryController),
		Node: new(node.NodeController),
		template: new(template.TemplateController),
	}

	iris.API("/", *r.Api)
	iris.API("/api/v1/history", *r.History)
	iris.API("/api/v1/node", *r.Node)
	iris.API("/api/v1/template", *r.template)
}
