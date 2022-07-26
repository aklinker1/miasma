package utils

// Returns fallback and the err if it exists, otherwise it returns the value and nil
func SafeReturn[T any](value T, fallback T, err error) (T, error) {
	if err != nil {
		return fallback, err
	} else {
		return value, err
	}
}
