package main

import (
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
