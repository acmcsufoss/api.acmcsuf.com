package mapper

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
)




// --- position ---
func PositionDtoToDomain(p *dto.Position) domain.Position {
	if p == nil {
		return domain.Position{}
	}

	return domain.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

func UpdatePositionDtoToDomain(p *dto.UpdatePosition) domain.UpdatePosition {
	if p == nil {
		return domain.UpdatePosition{}
	}

	return domain.UpdatePosition{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

// --- tier ---
func TierDtoToDomain(t *dto.Tier) domain.Tier {
	if t == nil {
		return domain.Tier{}
	}

	return domain.Tier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}

func UpdateTierDtoToDomain(t *dto.UpdateTier) domain.UpdateTier {
	if t == nil {
		return domain.UpdateTier{}
	}

	return domain.UpdateTier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}
