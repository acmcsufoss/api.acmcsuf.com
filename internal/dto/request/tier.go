package domain

type Tier struct {
	Tier   int    `json:"tier"`
	Title  string `json:"title"`
	Tindex int    `json:"t_index"`
	Team   string `json:"team"`
}
