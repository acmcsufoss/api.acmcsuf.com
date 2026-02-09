package dto_request

import "github.com/acmcsufoss/api.acmcsuf.com/internal/domain"

type Tier struct {
	Tier   int    `json:"tier"`
	Title  string `json:"title"`
	Tindex int    `json:"t_index"`
	Team   string `json:"team"`
}

func (t *Tier) ToDomain() domain.Tier {
	return domain.Tier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}
