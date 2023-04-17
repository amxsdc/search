package model

type Record struct {
	//gorm.Model
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
