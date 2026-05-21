package dto

type ErrorCode string

type Error struct {
	Code    ErrorCode `json:"code"`              // machine readable for user facing apps (e.g. 'EVENT_NOT_FOUND')
	Message string    `json:"message"`           // developer facing
	Details *any      `json:"details,omitempty"` // optional field or additional info
}

const (
	ErrOfficerNotFound  ErrorCode = "OFFICER_NOT_FOUND"
	ErrTierNotFound     ErrorCode = "TIER_NOT_FOUND"
	ErrPositionNotFound ErrorCode = "POSITION_NOT_FOUND"
	ErrBadRequestBody   ErrorCode = "INVALID_RQUEST_BODY"
	ErrUnknown          ErrorCode = "UNKNOWN_SERVER_ERROR"
)
