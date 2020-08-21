package models

import (
	"collectbackend/databases"
	"time"
)

type Category struct {
	ID        uint
	Name      string `gorm:"not null;unique"`
	DeletedAt *time.Time
	Subclass  uint // 0为最高级；
	Indxs     []Indx
}

func FindAllCategory() (catelist []Category, err error) {
	if err = databases.DB.Find(&catelist).Error; err != nil {
		return
	}
	return
}

func (cateo *Category) FindOne(id uint) (err error) {
	if err = databases.DB.Find(&cateo, "id=?", id).Error; err != nil {
		return
	}
	return
}

func (cateo *Category) Createnewone() uint {
	cateo.Name = "test"
	cateo.Subclass = 1

	databases.DB.Create(&cateo)
	return cateo.ID
}
