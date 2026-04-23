package forms

func NonEmptyPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func NonNilStr(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
