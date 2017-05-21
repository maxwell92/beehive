package mysql

import (
	"github.com/maxwell92/gokits/mysql"
)

type MysqlCore interface {
	Insert() error
	Query() error
	Delete() error
	Update() error
}

