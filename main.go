package main

import (
	"collectbackend/Router"
)

func main() {
	defer Mysql.DB.Close()
	Router.InitRouter()
}
