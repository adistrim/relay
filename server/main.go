package main

import (
	"relay/config"
	"relay/routes"
)

func main() {
	router := routes.MainRouter()
	router.Run(":" + config.ENV.Port)
}
