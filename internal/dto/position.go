package dto

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type Position struct {
	OfficerID string  `json:"officer_id"`
	Semester  string  `json:"semester"`
	Tier      int64   `json:"tier"`
	FullName  string  `json:"full_name"`
	Title     *string `json:"title"`
	Team      *string `json:"team"`
}

func PositionDomainToDto(p *domain.Position) Position {
	return Position{
		OfficerID: p.OfficerID,
		Semester:  p.Semester,
		Tier:      p.Tier,
		FullName:  p.FullName,
		Title:     p.Title,
		Team:      p.Team,
	}
}

func (p Position) ToDomain() domain.Position {
	return domain.Position{
		OfficerID: p.OfficerID,
		Semester:  p.Semester,
		Tier:      p.Tier,
		FullName:  p.FullName,
		Title:     p.Title,
		Team:      p.Team,
	}
}

type UpdatePosition struct {
	OfficerID string  `json:"officer_id"`
	Semester  string  `json:"semester"`
	Tier      int64   `json:"tier"`
	FullName  string  `json:"full_name"`
	Title     *string `json:"title"`
	Team      *string `json:"team"`
}

func (p UpdatePosition) ToDomain() domain.UpdatePosition {
	return domain.UpdatePosition{
		OfficerID: p.OfficerID,
		Semester:  p.Semester,
		Tier:      p.Tier,
		FullName:  p.FullName,
		Title:     p.Title,
		Team:      p.Team,
	}
}

type DeletePosition struct {
	OfficerID string `json:"officer_id"`
	Semester  string `json:"semester"`
	Tier      int64  `json:"tier"`
}

func (p DeletePosition) ToDomain() domain.DeletePosition {
	return domain.DeletePosition{
		OfficerID: p.OfficerID,
		Semester:  p.Semester,
		Tier:      p.Tier,
	}
}
