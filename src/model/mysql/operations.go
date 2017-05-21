package mysql

type OperationLogInfo interface {
	MysqlCore
}

type OperationLog struct {
}

func (o *OperationLog) Insert() error {

}

func (o *OperationLog) Query() error {

}

func (o *OperationLog) Update() error {

}

func (o *OperationLog) Delete() error {

}
