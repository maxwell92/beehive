package controller

import (
	bherror "common/error"
	"github.com/kataras/iris"
	"github.com/maxwell92/gokits/sorter"
	bhlog "github.com/maxwell92/gokits/log"
)

var log = bhlog.Log

type Controller struct {
	*iris.Context
	Be     *bherror.BeehiveError
	Sorter *sorter.Sorter
}

type IController interface {
	WriteError()
	ValidateSession(sessionId, orgId string)
	CheckError() bool // if error true, else false
	WriteOk()
	OrderedBy(less ...sorter.LessFunc) *sorter.Sorter
}

func (c *Controller) OrderedBy(less ...sorter.LessFunc) *sorter.Sorter {
	return &sorter.Sorter{
		Lessf: less,
	}
}

func (c *Controller) WriteError() {
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	log.Errorf("Controller Response BeehiveError: controller=%p, code=%d, msg=%s", c, c.Be.Code, bherror.Errors[c.Be.Code].LogMsg)
	c.Response.Header.Set("Cache-Control", "no-store")
	c.Write(c.Be.String())
}

func (c *Controller) CheckError() bool {
	if c.Be != nil {
		c.WriteError()
		return true
	}
	return false
}

func (c *Controller) WriteOk(data interface{}) {
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Be = bherror.NewBeehiveError(bherror.EOK, data)
	c.Response.Header.Set("Cache-Control", "no-store")
	c.Write(c.Be.String())
}
