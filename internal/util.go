package internal

func Map[T any](data []T, fn func(T) T) []T {
	mapped := make([]T, len(data))
	for i := range data {
		mapped[i] = fn(data[i])
	}

	return mapped
}
