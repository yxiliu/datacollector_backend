package services

import (
	"collectbackend/models"
)

type NewIndex struct {
	Name       string `json:"name"`
	Unit       string `json:"unit"`
	Memo       string `json:"memo"`
	Category   uint   `json:"cateid"`
	MonthC     uint8  `json:"month"`
	SeasonC    uint8  `json:"season"`
	Yearc      bool   `json:"year"`
	Maintainer string `json:"maintainer"`
}

// Insert 添加一个新的Index
func (ths *NewIndex) Insert() (id uint, err error) {
	var indx models.Indx
	id, err = indx.Insert(ths.Name, ths.Maintainer, ths.Unit, ths.Memo, ths.MonthC, ths.SeasonC, ths.Yearc, ths.Category)
	if err != nil {
		return
	}
	return
}
