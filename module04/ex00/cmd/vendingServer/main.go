package main

import (
	"fmt"

	router "github.com/Arclight-V/Golang/tree/module04-ex00/module04/ex00/internal/app/router"
)

func main() {
	r := router.NewRouter()

	r.Start()

	fmt.Println("Hello")
}
