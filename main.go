package main

import (
	databases "collectbackend/Databases"
	"collectbackend/router"
)

func main() {
	defer databases.DB.Close()
	router.InitRouter()
}
