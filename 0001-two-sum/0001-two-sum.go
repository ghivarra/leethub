func twoSum(nums []int, target int) []int {

    res := []int{}

	for i, num := range nums {
		diff := target - num

		for x, part := range nums {

			if i == x {
				continue
			}

			if part == diff {
				res = append(res, i, x)
				break
			}
		}

		if len(res) == 2 {
			break
		}
	}

	return res
}