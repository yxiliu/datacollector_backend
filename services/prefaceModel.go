package services

import (
	"collectbackend/models"
)

// Preface  最开始的主页
type Preface struct {
	CategoryID uint   `json:"categoryid"`
	Name       string `json:"name"`
}

// SearchAll 查找所有
func (pr *Preface) SearchAll() []models.Indxs {
	var cate models.Category
	cate.FindOne(pr.CategoryID)
	var indx models.Indxs
	return indx.FindAll(cate, pr.Name)
}
