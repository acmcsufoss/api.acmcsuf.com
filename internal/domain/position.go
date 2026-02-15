package domain

type Position struct {
	Oid      string
	Semester string
	Tier     int
	FullName string
	Title    *string
	Team     *string
}

type UpdatePosition struct {
	Oid      string
	Semester string
	Tier     int
	FullName string
	Title    *string
	Team     *string
}
