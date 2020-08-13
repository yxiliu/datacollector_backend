package main

import (
	"collectbackend/databases"
	"collectbackend/models"
	"collectbackend/router"
)

func main() {
	databases.DB.AutoMigrate(&models.Category{})
	databases.DB.AutoMigrate(&models.Indxs{})
	databases.DB.AutoMigrate(&models.IndxRecord{})
	defer databases.DB.Close()
	router.InitRouter()
}
