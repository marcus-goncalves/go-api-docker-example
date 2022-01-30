package main

import "go-docker-example/v1/routes"

func main() {
	router := routes.CreateRouter()

	router.Run(":3000")
}
