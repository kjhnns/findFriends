package main

import (
	"fmt"
	ctrl "github.com/kjhnns/findFriends/app"
)

func initRouting() {

	fmt.Println("Routing")
	g.Static("/public", "./public")
	g.LoadHTMLGlob("templates/*")
	g.StaticFile("/favicon.ico", "./public/images/favicon.png")

	g.GET("/", ctrl.Home)
	// g.GET("/", ctrl.NotFound)
}
