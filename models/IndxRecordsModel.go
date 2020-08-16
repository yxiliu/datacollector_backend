package models

import "collectbackend/databases"

type IndxRecord struct {
	ID       uint
	IndxsID  uint
	Maintype uint8 //0为Month；1为season；2为year
	Subtype  uint8 //0为average; 1为accumulation
	Amount   float64
	Year     uint
	Month    uint
}

func (idr *IndxRecord) Insert(indxId uint, maintype uint8, subtye uint8, amount float64, month uint, year uint) (id uint, err error) {
	newrecord := IndxRecord{
		IndxsID:  indxId,
		Maintype: maintype,
		Subtype:  subtye,
		Amount:   amount,
		Month:    month,
		Year:     year,
	}
	result := databases.DB.Create(&newrecord)
	id = idr.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func FindAllRecordByIndexAndYear(indxs uint, year uint) (listofindex []IndxRecord) {
	databases.DB.Where("year = ? and IndxsID = ?", year, indxs).Find(&listofindex)
	return listofindex
}

func (idr *IndxRecord) UpdateCurrent(id uint, amount float64) {
	var indxr IndxRecord
	databases.DB.First(&indxr, "id=?", id)
	indxr.Amount = amount
	databases.DB.Save(&indxr)
}
