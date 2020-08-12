package Models

import (
	"collectbackend/Databases"
	"time"
)

type Indxs struct {
	Id          uint
	Name        string `gorm:"not null;unique"`
	Maintainer  string
	Memo        string `gorm:"size:255"`
	Categories  Category
	IndxRecords []IndxRecord
	MonthIndx   uint8 //0为不选；1为单选均；2为单选累计；3为全选
	SeasonIndx  uint8
	YearIndx    bool // 年只有年均
	idxUnit     string
	DeletedAt   *time.Time
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// // }

func (this *Indxs) Insert(name string, Maintainer string, idxUnit string, memo string, MonthIndx uint8, seasonIndx uint8, yearIndx bool, category string) (id uint, err error) {
	result := Databases.DB.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (this *Indxs) FindAll(category Category, Name string) (listofindex []Indxs) {
	if Name != "" {
		Databases.DB.Where("name = ?", Name).Find(&listofindex)
	} else if category != (Category{}) {
		Databases.DB.Where("categories = ?", category).Find(&listofindex)
	}
	return
}

func (this *Indxs) Del(Id uint) {
	var indx Indxs
	Databases.DB.First(&indx, "id = ?", Id)
	Databases.DB.Delete(&indx)
}

func (this *Indxs) FindOne(Id uint) {
	Databases.DB.First(&this, "id = ?", Id)
}
