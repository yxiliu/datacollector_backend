package Services

import (
	"collectbackend/Models"
)

type IndxInfo struct {
	Id      int    `json:"id"`
	Testcol string `json:"testcol"`
}

func (this *IndxInfo) GetInfoDetail() Models.Indxs {
	var indxinfo Models.Indxs
	indxinfo.
		testModel.Id = this.Id
	testModel.Testcol = this.Testcol
	id, err = testModel.Insert()
	return
}
