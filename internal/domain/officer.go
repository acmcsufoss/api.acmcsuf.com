package domain

type Officer struct {
	Uuid     string
	FullName string
	Picture  *string
	Github   *string
	Discord  *string
}

type UpdateOfficer struct {
	FullName *string
	Picture  *string
	Github   *string
	Discord  *string
}
