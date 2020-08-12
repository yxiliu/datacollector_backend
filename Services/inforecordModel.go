package Services

import (
	"collectbackend/Models"
)

type IndxInfo struct {
	Id int `json:"id"`
}

func (this *Test) GetInfo() (id int, err error) {
	var testModel Models.Test
	testModel.Id = this.Id
	testModel.Testcol = this.Testcol
	id, err = testModel.Insert()
	return
}
