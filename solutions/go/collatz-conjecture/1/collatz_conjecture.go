package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return n, errors.New("non pos error")
	}
	steps := 0
	for {
		switch {
		case n == 1:
			return steps, nil
		case n%2 == 0:
			n = n / 2
		default:
			n = n*3 + 1
		}
		steps++
	}
}
