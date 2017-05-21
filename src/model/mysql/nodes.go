package mysql

import (

)


type NodeInfo interface {
	MysqlCore

}

type Node struct {

}

func (n *Node) Insert() error {

}

func (n *Node) Query() error {

}

func (n *Node) Update() error {

}

func (n *Node) Delete() error {

}

