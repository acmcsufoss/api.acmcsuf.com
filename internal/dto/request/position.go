package dto_request

type Position struct {
	Oid      string `json:"oid"`
	Semester string `json:"semester"`
	Tier     int    `json:"tier"`
}
