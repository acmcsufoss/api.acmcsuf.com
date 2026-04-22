package dto

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type Officer struct {
	Uuid     string  `json:"uuid"`
	FullName string  `json:"full_name"`
	Picture  *string `json:"picture"`
	Github   *string `json:"github"`
	Discord  *string `json:"discord"`
}

func (o *Officer) ToDomain() domain.Officer {
	if o == nil {
		return domain.Officer{}
	}

	return domain.Officer{
		Uuid:     o.Uuid,
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}

func OfficerDomainToDto(o domain.Officer) Officer {
	return Officer{
		Uuid:     o.Uuid,
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}

type UpdateOfficer struct {
	Uuid     string  `json:"uuid"`
	FullName *string `json:"full_name"`
	Picture  *string `json:"picture"`
	Github   *string `json:"github"`
	Discord  *string `json:"discord"`
}

func (o *UpdateOfficer) ToDomain() domain.UpdateOfficer {
	if o == nil {
		return domain.UpdateOfficer{}
	}

	return domain.UpdateOfficer{
		Uuid:     o.Uuid,
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}
