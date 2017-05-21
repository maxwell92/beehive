package main

import (
	"common/router"
	"common/server"
	"common/config"
	"github.com/maxwell92/gokits/mysql"
	"fmt"
)

func init() {
	router.Setup()
	config.Instance().Load()
	mysql.NewMysqlClient(config.Instance().DbHost, config.Instance().DbPort, config.Instance().DbUser, config.Instance().DbPass, config.Instance().DbName, config.Instance().DbPool).Open()
}

func main() {
	fmt.Println("Hello World")
	server.Instance().Up(":8080")
}
