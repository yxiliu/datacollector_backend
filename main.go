package main

import (
	"collectbackend/databases"
	"collectbackend/router"
)

func main() {
	//databases.DB.AutoMigrate(&models.Category{})
	//databases.DB.AutoMigrate(&models.Indx{})
	//databases.DB.AutoMigrate(&models.IndxRecord{})
	defer databases.DB.Close()
	router.InitRouter()
}
