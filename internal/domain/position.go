package domain

import dto_response "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/response"

type Position struct {
	Oid      string
	Semester string
	Tier     int
}

func (p *Position) ToDTO() dto_response.Position {
	return dto_response.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
	}
}
