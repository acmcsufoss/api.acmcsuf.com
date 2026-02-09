package domain

import dto_response "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/response"

type Tier struct {
	Tier   int
	Title  string
	Tindex int
	Team   string
}

func (t *Tier) ToDTO() dto_response.Tier {
	return dto_response.Tier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}
