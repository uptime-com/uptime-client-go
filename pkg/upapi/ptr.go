package upapi

// BoolPtr returns a pointer to the given bool value.
// Use this helper to set *bool fields on Check structs for PATCH operations.
func BoolPtr(v bool) *bool {
	return &v
}

func StringPtr(v string) *string {
	return &v
}
