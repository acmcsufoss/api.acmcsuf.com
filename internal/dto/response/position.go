package dto_response

type Position struct {
	Oid      string `json:"oid,omitempty"`
	Semester string `json:"semester"`
	Tier     int    `json:"tier"`
}
