package main

import (
	"LoanTrackerAPI/Delivery/routers"
	infrastructure "LoanTrackerAPI/Infrastructure"
	"fmt"

	xss "github.com/araujo88/gin-gonic-xss-middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	confi := cors.DefaultConfig()
	confi.AllowAllOrigins = true

	server.Use(cors.New(confi))
	var xssMdlwr xss.XssMw
	server.Use(xssMdlwr.RemoveXss())

	config, err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	fmt.Print(config)
	DB, client, err := infrastructure.ConnectDB(config.DatabaseUrl, config.Dbname)

	if err != nil {
		fmt.Print("Error in connectDB")
	}
	defer infrastructure.CloseDB(client)
	routers.Router(server, *config, DB)
	server.Run(fmt.Sprintf(":%d", config.Port))

}
