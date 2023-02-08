package main

import (
	"github.com/caibo86/ginblog/model"
	"github.com/caibo86/ginblog/routes"
	"github.com/caibo86/ginblog/utils"
	"github.com/caibo86/ginblog/utils/errmsg"
	"log"
)

func main() {

	username := "caibo"
	ss, err := utils.SetToken(username)
	if err != errmsg.OK {
		log.Println(err)
		return
	}
	log.Println(ss)
	claims, err := utils.CheckToken(ss)
	if err != errmsg.OK {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", claims)
	model.InitDb()
	routes.InitRouter()
}
