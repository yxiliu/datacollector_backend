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

func (nre *NewRecord) Insert() (id uint, err error) {
	var record models.IndxRecord

	id, err = record.Insert(nre.Indxid, nre.Maintype, nre.Subtype, nre.Amount, nre.Month, nre.Year)
	if err != nil {
		return
	}
	return
}
