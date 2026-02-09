package domain

type Event struct {
	Location string `json:"location"`
	StartAt  int64  `json:"start_at"`
	EndAt    int64  `json:"end_at"`
	IsAllDay bool   `json:"is_all_day"`
	Host     string `json:"host"`
}
