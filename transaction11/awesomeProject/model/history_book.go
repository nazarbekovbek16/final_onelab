package model

type HistoryBook struct {
	ID       int `gorm:"primaryKey" json:"id"`
	UserID   int `json:"user_id"`
	BookID   int
	Duration int `json:"duration"`
	Price    int
	IsPaid   bool
	IsGiven  bool
}
