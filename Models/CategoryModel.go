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

func (this *Category) FindAll() ([]*Category, error) {
	var catelist []*Category
	err := Databases.DB.Find(&catelist).Error
	if err != nil {
		return nil, err
	}
	return catelist, nil
}
