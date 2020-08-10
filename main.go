package main

import (
	"gin/Router"
)

func main() {
	defer Mysql.DB.Close()
	Router.InitRouter()
}
