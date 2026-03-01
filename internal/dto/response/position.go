package dto_response

type Position struct {
	Oid      string  `json:"oid"`
	Semester string  `json:"semester"`
	Tier     int     `json:"tier"`
	FullName string  `json:"tier"`
	Title    *string `json:"title"`
	Team     *string `json:"team"`
}
