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

func (this *NewIndex) Insert() (id uint, err error) {
	var cate models.Category
	cate.FindOne(this.Category)
	var indx models.Indxs
	id, err = indx.Insert(this.Name, this.Maintainer, this.Unit, this.Memo, this.MonthC, this.SeasonC, this.Yearc, cate)
	if err != nil {
		return
	}
	return
}
