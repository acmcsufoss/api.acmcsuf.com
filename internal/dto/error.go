package dto

type Error struct {
	code    string // machine readable for user facing apps (e.g. 'EVENT_NOT_FOUND')
	message string // developer facing
	details *any   // optional field or additional info
}
