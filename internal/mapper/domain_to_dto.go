package mapper

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
)

// --- postition ---
func PositionDomainToDto(p *domain.Position) dto.Position {
	return dto.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

// --- tier ---
func TierDomainToDto(t *domain.Tier) dto.Tier {
	return dto.Tier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}
