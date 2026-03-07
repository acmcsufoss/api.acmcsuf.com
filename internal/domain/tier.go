package domain

type Tier struct {
	Tier   int
	Title  *string
	Tindex *int
	Team   *string
}

type UpdateTier struct {
	Tier   int
	Title  *string
	Tindex *int
	Team   *string
}
