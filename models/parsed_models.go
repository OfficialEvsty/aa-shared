package models

// Item json struct holds data
type Item struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Tier     int    `json:"tier"`
	TierURL  string `json:"tierUrl"`
	ImageURL string `json:"image_url"`
}

// DropItem json struct holds data of boss's drop items
type DropItem struct {
	Item     Item   `json:"item"`
	Quantity string `json:"quantity" yaml:"rate"`
}

// Boss struct represents parsed data from AA db site about boss and his loot
type Boss struct {
	ID   int         `json:"boss_id"`
	Name string      `json:"boss_name"`
	Loot []*DropItem `json:"loot"`
}

// Server parsed data
type Server struct {
	ID int `json:"server_id"`
}

// Unit data about each server
type Unit struct {
	ID int `json:"unit_id"`
}
