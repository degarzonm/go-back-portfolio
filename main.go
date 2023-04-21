package main

import (
    "os"

    middleware "github.com/degarzonm/go-back-portfolio/middleware"
    routes "github.com/degarzonm/go-back-portfolio/routes"

    "github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
    

    router := gin.New()
    router.Use(gin.Logger())
    routes.UserRegistryRoutes(router)

    router.Use(middleware.Authentication())

	routes.SoldierRoutes(router)
	routes.OfficerRoutes(router)
    routes.GeneralRoutes(router)
	routes.RecruiterRoutes(router)

    router.Run(os.Getenv("PORT"))
}
