package Models

import (
	"collectbackend/Databases"
	"time"
)

type Category struct {
	Id        uint
	Name      string `gorm:"not null;unique"`
	DeletedAt *time.Time
	subclass  uint // 0为最高级；
}

func (this *Category) FindAll() (catelist []Category, err error) {
	if err = Databases.DB.Find(&catelist).Error; err != nil {
		return
	}
	return
}

func (this *Category) FindOne(id uint) (err error) {
	if err = Databases.DB.Find(&this, "id=?", id).Error; err != nil {
		return
	}
	return
}
