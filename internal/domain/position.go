package domain

type Position struct {
	OfficerID string
	Semester  string
	Tier      int64
	FullName  string
	Title     *string
	Team      *string
}

type UpdatePosition struct {
	OfficerID string
	Semester  string
	Tier      int64
	FullName  string
	Title     *string
	Team      *string
}

type DeletePosition struct {
	OfficerID string
	Semester  string
	Tier      int64
}

type GetPosition DeletePosition
