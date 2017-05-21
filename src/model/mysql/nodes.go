package mysql

import (

)


type Node interface {
	Insert() error
	Query() error
}

type NodeInfo struct {

}

func (ni *NodeInfo)Insert() error {

}

func (ni *NodeInfo)Query() error {

}


