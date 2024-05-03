package assert

import "log"

func Equal[T comparable](expected, actual T) {
	if expected != actual {
		log.Fatalf("value is not equal. expected: %v. actual: %v.", expected, actual)
	}
}

func NotEqual[T comparable](expected, actual T) {
	if expected == actual {
		log.Fatalf("value is equal. expected: %v. actual: %v.", expected, actual)
	}
}

func NotEmpty[T any](arr []T) {
	if len(arr) == 0 {
		log.Fatalf("provided array is empty %v", arr)
	}
}
