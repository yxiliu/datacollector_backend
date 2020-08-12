package Models

import "collectbackend/Databases"

type IndxRecord struct {
	Id       uint
	Indxs    Indxs
	maintype uint8 //0为Month；1为season；2为year
	subtype  uint8 //0为average; 1为accumulation
	amount   float64
	year     uint
	month    uint
}

func (this *IndxRecord) Insert(indx Indxs, maintype uint8, subtye uint8, amount float64) (id uint, err error) {
	newrecord := IndxRecord{
		Indxs:    indx,
		maintype: maintype,
		subtype:  subtye,
		amount:   amount,
	}
	result := Databases.DB.Create(&newrecord)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (this *IndxRecord) FindAll(indxs Indxs, year uint) (listofindex []IndxRecord) {
	Databases.DB.Where("year = ? and Indxs = ?", year, indxs).Find(&listofindex)
	return listofindex
}

func (this *IndxRecord) UpdateCurrent(id uint, amount float64) {
	var indxr IndxRecord
	Databases.DB.First(&indxr, "id=?", id)
	indxr.amount = amount
	Databases.DB.Save(&indxr)
}
