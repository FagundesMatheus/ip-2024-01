func ISort(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		j := i                             
		for j > 0 && nums[j] < nums[j-1] {
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j-- 
		}
	}
	return nums
}