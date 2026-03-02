package dto_request

type TierDto struct {
	Tier   int     `json:"tier"`
	Title  *string `json:"title"`
	Tindex *int    `json:"t_index"`
	Team   *string `json:"team"`
}

type UpdateTier struct {
	Tier   int     `json:"tier"`
	Title  *string `json:"title"`
	Tindex *int    `json:"t_index"`
	Team   *string `json:"team"`
}
