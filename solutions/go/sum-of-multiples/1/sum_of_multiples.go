package summultiples

// SumMultiples brute force.
func SumMultiples(limit int, divisors ...int) int {
	numbers := map[int]bool{}
	for i := 0; i < len(divisors); i++ {
		if divisors[i] == 0 {
			numbers[0] = true
			continue
		}
		for j := divisors[i]; j < limit; j += divisors[i] {
			numbers[j] = true
		}
	}
	sum := 0
	for k, _ := range numbers {
		sum += k
	}
	return sum
}
