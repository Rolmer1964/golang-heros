package model

type HeroPowerStatsResponse struct {
	Response     string `json:"response"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}
