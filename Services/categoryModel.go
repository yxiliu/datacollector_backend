package Services

import (
	"collectbackend/Models"
)

type Test struct {
	Id      uint    `json:"id"`
}

func (this *Test) Insert() (id int, err error) {
	var testModel Models.Test
	testModel.Id = this.Id
	testModel.Testcol = this.Testcol
	id, err = testModel.Insert()
	return
}
