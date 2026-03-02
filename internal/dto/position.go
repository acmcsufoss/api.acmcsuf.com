package dto_request

type PositionDto struct {
	Oid      string  `json:"oid"`
	Semester string  `json:"semester"`
	Tier     int     `json:"tier"`
	FullName string  `json:"full_name"`
	Title    *string `json:"title"`
	Team     *string `json:"team"`
}

type UpdatePosition struct {
	Oid      string  `json:"oid"`
	Semester string  `json:"semester"`
	Tier     int     `json:"tier"`
	FullName string  `json:"full_name"`
	Title    *string `json:"title"`
	Team     *string `json:"team"`
}
