package internal

import (
	"strconv"
)

func Map[T any](data []T, fn func(T) T) []T {
	mapped := make([]T, len(data))
	for i := range data {
		mapped[i] = fn(data[i])
	}

	return mapped
}

func ParseInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func SliceLast[T any](data []T) T {
	return data[len(data)-1]
}
