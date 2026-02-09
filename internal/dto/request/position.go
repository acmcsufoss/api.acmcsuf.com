package dto_request

import "github.com/acmcsufoss/api.acmcsuf.com/internal/domain"

type Position struct {
	Oid      string `json:"oid"`
	Semester string `json:"semester"`
	Tier     int    `json:"tier"`
}

func (p *Position) ToDomain() domain.Position {
	return domain.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
	}
}
