package services

import "collectbackend/models"

// IndxInfo 给一个ID，给具体指标的属性
type IndxInfo struct {
	ID uint `json:"id"`
}

// GetEachInfo 找一个
func (ifo *IndxInfo) GetEachInfo() (indxinfo models.Indxs) {
	indxinfo.FindOne(ifo.ID)
	return
}
