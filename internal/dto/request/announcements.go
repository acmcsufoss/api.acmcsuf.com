package domain

type Announcement struct {
	Visibility       string `json:"visibility"`
	AnnounceAt       int64  `json:"announce_at"`
	DiscordChannelID string `json:"discord_channel_id"`
	DiscordMessageID string `json:"discord_message_id"`
}
