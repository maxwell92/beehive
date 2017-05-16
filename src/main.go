package main

import (
	"fmt"
	"common/router"
	"common/server"
)

func init() {
	router.Setup()
}

func main() {
	fmt.Println("Hello World")
	server.Instance().Up(":8080")
}
