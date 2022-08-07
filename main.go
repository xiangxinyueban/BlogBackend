package main

import (
	"blogbackend/conf"
	"blogbackend/models"
	"blogbackend/router"
)

func main() {
	conf.Init()
	models.Init()
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
