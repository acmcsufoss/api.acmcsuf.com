package forms

func NonEmptyPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
