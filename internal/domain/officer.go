package domain

import dto_response "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/response"

type Officer struct {
	Uuid     string
	FullName string
	Picture  string
	Github   string
	Discord  string
}

func (o *Officer) ToDTO() dto_response.Officer {
	return dto_response.Officer{
		Uuid:     o.Uuid,
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}
