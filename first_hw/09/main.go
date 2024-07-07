func arrSum(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	return result
}
