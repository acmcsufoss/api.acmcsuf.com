package dto_request

type OfficerDto struct {
	Uuid     string  `json:"uuid"`
	FullName string  `json:"full_name"`
	Picture  *string `json:"picture"`
	Github   *string `json:"github"`
	Discord  *string `json:"discord"`
}

type UpdateOfficer struct {
	Uuid     string  `json:"uuid"`
	FullName *string `json:"full_name"`
	Picture  *string `json:"picture"`
	Github   *string `json:"github"`
	Discord  *string `json:"discord"`
}
