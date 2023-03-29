package main

import (
	"go-todo/config"
	"go-todo/routers"
)

func main() {
	router := routers.InitRoute()
	port := config.EnvVar("SERVER_PORT", ":8080")
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}
