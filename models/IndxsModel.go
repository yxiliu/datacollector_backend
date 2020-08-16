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

// Insert 添加一个
func (idx *Indx) Insert(name string, Maintainer string, idxUnit string, memo string, MonthIndx uint8, seasonIndx uint8, yearIndx bool, category uint) (id uint, err error) {
	idx.Name = name
	idx.Maintainer = Maintainer
	idx.IdxUnit = idxUnit
	idx.Memo = memo
	idx.MonthIndx = MonthIndx
	idx.SeasonIndx = seasonIndx
	idx.YearIndx = yearIndx
	idx.CategoryID = category
	result := databases.DB.Create(&idx)
	id = idx.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// FindAllByCate 通过category找到所有的
func FindAllByCate(category uint) (listofindex []Indx) {
	databases.DB.Where("categories = ?", category).Find(&listofindex)
	return
}

// FindAllByName 通过Name找到所有的 搜索的时候用
func FindAllByName(Name string) (listofindex []Indx) {
	databases.DB.Where("name = ?", Name).Find(&listofindex)
	return
}

// FindAllIndx 没有任何条件的
func FindAllIndx() (listofindex []Indx) {
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
