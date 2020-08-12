package Services

import (
	"collectbackend/Models"
)

type Preface struct {
	categoryId uint   `json:"categoryid"`
	name       string `json:"name"`
}

func (this *Preface) SearchAll() []Models.Indxs {
	var cate Models.Category
	cate.FindOne(this.categoryId)
	var indx Models.Indxs
	return indx.FindAll(cate, this.name)
}
