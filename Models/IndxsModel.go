package models

import (
	databases "collectbackend/databases"
	"time"
)

type Indxs struct {
	ID          uint
	Name        string `gorm:"not null;unique"`
	Maintainer  string
	Memo        string `gorm:"size:255"`
	Categories  Category
	IndxRecords []IndxRecord
	MonthIndx   uint8 //0为不选；1为单选均；2为单选累计；3为全选
	SeasonIndx  uint8
	YearIndx    bool // 年只有年均
	IdxUnit     string
	DeletedAt   *time.Time
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// // }

func (idx *Indxs) Insert(name string, Maintainer string, idxUnit string, memo string, MonthIndx uint8, seasonIndx uint8, yearIndx bool, category Category) (id uint, err error) {
	result := databases.DB.Create(&idx)
	id = idx.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (idx *Indxs) FindAll(category Category, Name string) (listofindex []Indxs) {
	if Name != "" {
		databases.DB.Where("name = ?", Name).Find(&listofindex)
	} else if category != (Category{}) {
		databases.DB.Where("categories = ?", category).Find(&listofindex)
	}
	return
}

func (idx *Indxs) Del(ID uint) {
	var indx Indxs
	databases.DB.First(&indx, "id = ?", ID)
	databases.DB.Delete(&indx)
}

func (idx *Indxs) FindOne(ID uint) {
	databases.DB.First(&idx, "id = ?", ID)
}
