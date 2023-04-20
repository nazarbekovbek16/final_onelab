package model

type Card struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	UserID uint    `json:"userID"`
	Money  float64 `json:"money"`
}
