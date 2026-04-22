package store

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

func AnnouncementDomainToDB(announcement domain.Announcement) dbmodels.CreateAnnouncementParams {
	return dbmodels.CreateAnnouncementParams{
		Uuid:             announcement.Uuid,
		Visibility:       announcement.Visibility,
		AnnounceAt:       announcement.AnnounceAt.Unix(),
		DiscordChannelID: stringToNullString(announcement.DiscordChannelID),
		DiscordMessageID: stringToNullString(announcement.DiscordMessageID),
	}
}

func UpdateAnnouncementDomainToDB(announcement domain.UpdateAnnouncement) dbmodels.UpdateAnnouncementParams {
	return dbmodels.UpdateAnnouncementParams{
		Uuid:             announcement.Uuid,
		Visibility:       stringToNullString(announcement.Visibility),
		AnnounceAt:       timeToNullInt64(announcement.AnnounceAt),
		DiscordChannelID: stringToNullString(announcement.DiscordChannelID),
		DiscordMessageID: stringToNullString(announcement.DiscordMessageID),
	}
}

func AnnouncementDBToDomain(announcement dbmodels.Announcement) domain.Announcement {
	return domain.Announcement{
		Uuid:             announcement.Uuid,
		Visibility:       announcement.Visibility,
		AnnounceAt:       time.Unix(announcement.AnnounceAt, 0),
		DiscordChannelID: nullStringPtr(announcement.DiscordChannelID),
		DiscordMessageID: nullStringPtr(announcement.DiscordMessageID),
	}
}
