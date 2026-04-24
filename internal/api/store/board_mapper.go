package store

import "github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"
import "github.com/acmcsufoss/api.acmcsuf.com/internal/domain"

func OfficerDomainToDB(officer domain.Officer) dbmodels.CreateOfficerParams {
	return dbmodels.CreateOfficerParams{
		Uuid:     officer.Uuid,
		FullName: officer.FullName,
		Picture:  stringToNullString(officer.Picture),
		Github:   stringToNullString(officer.Github),
		Discord:  stringToNullString(officer.Discord),
	}
}

func UpdateOfficerDomainToDB(officer domain.UpdateOfficer) dbmodels.UpdateOfficerParams {
	return dbmodels.UpdateOfficerParams{
		FullName: stringValue(officer.FullName),
		Picture:  stringToNullString(officer.Picture),
		Github:   stringToNullString(officer.Github),
		Discord:  stringToNullString(officer.Discord),
	}
}

func OfficerDBToDomain(officer dbmodels.Officer) domain.Officer {
	return domain.Officer{
		Uuid:     officer.Uuid,
		FullName: officer.FullName,
		Picture:  nullStringPtr(officer.Picture),
		Github:   nullStringPtr(officer.Github),
		Discord:  nullStringPtr(officer.Discord),
	}
}

func GetOfficerDBToDomain(row dbmodels.GetOfficerRow) domain.Officer {
	return domain.Officer{
		FullName: row.FullName,
		Picture:  nullStringPtr(row.Picture),
		Github:   nullStringPtr(row.Github),
		Discord:  nullStringPtr(row.Discord),
	}
}

func TierDomainToDB(tier domain.Tier) dbmodels.CreateTierParams {
	return dbmodels.CreateTierParams{
		Tier:   int64(tier.Tier),
		Title:  stringToNullString(tier.Title),
		TIndex: intToNullInt64(tier.Tindex),
		Team:   stringToNullString(tier.Team),
	}
}

func UpdateTierDomainToDB(tier domain.UpdateTier) dbmodels.UpdateTierParams {
	return dbmodels.UpdateTierParams{
		Tier:   int64(tier.Tier),
		Title:  stringToNullString(tier.Title),
		TIndex: intToNullInt64(tier.Tindex),
		Team:   stringToNullString(tier.Team),
	}
}

func TierDBToDomain(tier dbmodels.Tier) domain.Tier {
	var tIndex *int
	if tier.TIndex.Valid {
		v := int(tier.TIndex.Int64)
		tIndex = &v
	}

	return domain.Tier{
		Tier:   int(tier.Tier),
		Title:  nullStringPtr(tier.Title),
		Tindex: tIndex,
		Team:   nullStringPtr(tier.Team),
	}
}

func PositionDomainToDB(position domain.Position) dbmodels.CreatePositionParams {
	return dbmodels.CreatePositionParams{
		OfficerID: position.OfficerID,
		Semester:  position.Semester,
		Tier:      int64(position.Tier),
		FullName:  position.FullName,
		Title:     stringToNullString(position.Title),
		Team:      stringToNullString(position.Team),
	}
}

func UpdatePositionDomainToDB(position domain.UpdatePosition) dbmodels.UpdatePositionParams {
	return dbmodels.UpdatePositionParams{
		OfficerID: position.OfficerID,
		Semester:  position.Semester,
		Tier:      int64(position.Tier),
		FullName:  position.FullName,
		Title:     stringToNullString(position.Title),
		Team:      stringToNullString(position.Team),
	}
}

func DeletePositionDomainToDB(position domain.DeletePosition) dbmodels.DeletePositionParams {
	return dbmodels.DeletePositionParams{
		OfficerID: position.OfficerID,
		Semester:  position.Semester,
		Tier:      position.Tier,
	}
}

func PositionDBToDomain(position dbmodels.Position) domain.Position {
	return domain.Position{
		OfficerID: position.OfficerID,
		Semester:  position.Semester,
		Tier:      position.Tier,
		FullName:  position.FullName,
		Title:     nullStringPtr(position.Title),
		Team:      nullStringPtr(position.Team),
	}
}
