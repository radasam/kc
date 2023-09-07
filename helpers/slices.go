package helpers

func FilterSlice[T comparable](slice []T, item T) []T {
	sliceOut := []T{}

	for _, v := range slice {
		if v != item {
			sliceOut = append(sliceOut, v)
		}
	}

	return sliceOut
}
