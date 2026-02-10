package dto_request

import "github.com/acmcsufoss/api.acmcsuf.com/internal/domain"

type Officer struct {
	FullName string `json:"full_name,omitempty"`
	Picture  string `json:"picture"`
	Github   string `json:"github"`
	Discord  string `json:"discord"`
}

func (o *Officer) ToDomain() domain.Officer {
	return domain.Officer{
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}
