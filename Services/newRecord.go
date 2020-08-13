package services

import (
	"collectbackend/models"
)

type NewRecord struct {
	Indxid   uint    `json:"indexid"`
	Year     uint    `json:"year"`
	Month    uint    `json:"month"`
	Maintype uint8   `json:"maintype"`
	Subtype  uint8   `json:"subtype"`
	Amount   float64 `json:"amount"`
}

func (this *NewRecord) Insert() (id uint, err error) {
	var record models.IndxRecord
	var indx models.Indxs
	indx.FindOne(this.Indxid)
	id, err = record.Insert(indx, this.Maintype, this.Subtype, this.Amount, this.Month, this.Year)
	if err != nil {
		return
	}
	return
}
