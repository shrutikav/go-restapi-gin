package models

type Learn_recipe_content struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	CategoryType string `gorm:"type:varchar(255)" json:"category_type"`
	Title        string `gorm:"type:varchar(255)" json:"title"`
	Desc1        string `gorm:"type:varchar(255)" json:"desc1"`
	Footer       string `gorm:"type:varchar(255)" json:"footer"`
}
