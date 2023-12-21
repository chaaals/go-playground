package getgrade

func sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return total
}