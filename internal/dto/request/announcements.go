package dto_request

type Announcement struct {
	Uuid             string `json:"uuid"`
	Visibility       string `json:"visibility"`
	AnnounceAt       int64  `json:"announce_at"`
	DiscordChannelID string `json:"discord_channel_id"`
	DiscordMessageID string `json:"discord_message_id"`
}
