package array

/*
func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
*/

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}
	return sums
}
