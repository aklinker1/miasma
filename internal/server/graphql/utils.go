package graphql

func safeReturn[T any](value T, err error) (T, error) {
	var nilAny any
	if err != nil {
		return nilAny.(T), err
	} else {
		return value, err
	}
}
