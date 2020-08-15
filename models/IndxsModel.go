package models

import (
	databases "collectbackend/databases"
	"time"
)

type Indx struct {
	ID          uint
	Name        string `gorm:"not null;unique"`
	Maintainer  string
	Memo        string `gorm:"size:255"`
	CategoryID  uint
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

func (idx *Indx) Insert(name string, Maintainer string, idxUnit string, memo string, MonthIndx uint8, seasonIndx uint8, yearIndx bool, category Category) (id uint, err error) {
	result := databases.DB.Create(&idx)
	id = idx.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// FindAll 通过category找到所有的
func FindAllByCate(category uint) (listofindex []Indx) {
	databases.DB.Where("categories = ?", category).Find(&listofindex)
	return
}

// FindAll 通过Name找到所有的 搜索的时候用
func FindAllByName(Name string) (listofindex []Indx) {
	databases.DB.Where("name = ?", Name).Find(&listofindex)
	return
}
// FindAll 没有任何条件的
func FindAllIndx() (listofindex []Indx)  {
	databases.DB.Find(&listofindex)
	return
}

func (idx *Indx) Del(ID uint) {
	var indx Indx
	databases.DB.First(&indx, "id = ?", ID)
	databases.DB.Delete(&indx)
}

func (idx *Indx) FindOne(ID uint) {
	databases.DB.First(&idx, "id = ?", ID)
}
