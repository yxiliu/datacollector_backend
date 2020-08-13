package services

import (
	"collectbackend/models"
)

type IndxRecordInfo struct {
	ID   uint `json:"id"`
	Year uint `json:"year"`
}

func (idxr *IndxRecordInfo) GetTable() (indxlist []models.IndxRecord) {
	var record models.IndxRecord
	var indx models.Indxs
	indx.FindOne(idxr.ID)
	indxlist = record.FindAll(indx, idxr.Year)
	return
}
