package main

import (
	"todo-mysql/config"
	"todo-mysql/routing"
)

func main()  {
	config.Connect()
	routing.Routing()
}