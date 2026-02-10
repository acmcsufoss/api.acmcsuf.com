package dto_response

type Officer struct {
	Uuid     string `json:"uuid,omitempty"`
	FullName string `json:"full_name"`
	Picture  string `json:"picture"`
	Github   string `json:"github"`
	Discord  string `json:"discord"`
}
