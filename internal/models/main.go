package models

type (
	Client struct {
		Name         string `json:"name"`
		Tin          string `gorm:"size:10; unique" json:"tin"`
		RateID       uint   `json:"rateId"`
		Discount     rune   `json:"discount"`
		IsCreditable bool   `json:"isCreditable"`
		HasBenefit   bool   `json:"hasBenefit"`
	}
)
