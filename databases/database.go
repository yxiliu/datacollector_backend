package databases

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB  数据库连接handler
var DB *gorm.DB

func init() {
	var err error
	//DB, err = gorm.Open("mysql", "dqb:123456Hb@tcp(10.18.154.202:3306)/argsdb?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("mysql", "jingjing1:jingjingbangbang@tcp(123.207.68.151:3306)/argsdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

}
