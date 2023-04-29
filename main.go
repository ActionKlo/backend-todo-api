package main

import "todoBackedAPI/routers"

func main() {
	router := routers.SetupRouters()

	router.Run("localhost:8080")
}
