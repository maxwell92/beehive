package main

import (
	"fmt"
	"common/server"
)

func init() {

}

func main() {
	fmt.Println("Hello World")
	server.Instance().Up(":8080")
}
