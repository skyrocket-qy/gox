package slice

func Contains[T comparable](slice []T, key T) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}
