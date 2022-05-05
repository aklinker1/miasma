package utils

func Int32Or(ptr *int32, defaultValue int32) int32 {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

func BoolOr(ptr *bool, defaultValue bool) bool {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}
