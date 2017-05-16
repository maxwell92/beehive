package main

import (
	"common/router"
	"common/server"
	"common/config"
	"fmt"
)

func init() {
	router.Setup()
	config.Instance().Load()
}

func main() {
	fmt.Println("Hello World")
	server.Instance().Up(":8080")
}
