func sumList(ref []int, nums []int) []int {
	var res []int

	addedNumber := 0
	sumNumber := 0

	for i, value := range ref {

		if (i + 1) <= len(nums) {
			sumNumber = nums[i]
		} else {
			sumNumber = 0
		}

		count := value + sumNumber + addedNumber

		if count > 9 {
			count = count % 10
			addedNumber = 1
		} else {
			addedNumber = 0
		}

		res = append([]int{count}, res...)

		// added new array index if there is more than one
		if (i+1) >= len(ref) && addedNumber == 1 {
			res = append([]int{1}, res...)
		}
	}

	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// extract value
	var list_1 []int
	for l1 != nil {
		list_1 = append(list_1, l1.Val)
		l1 = l1.Next
	}

	// extract value
	var list_2 []int
	for l2 != nil {
		list_2 = append(list_2, l2.Val)
		l2 = l2.Next
	}

	var sumResult []int
	if len(list_1) > len(list_2) {
		sumResult = sumList(list_1, list_2)
	} else {
		sumResult = sumList(list_2, list_1)
	}

	// add into struct
	var res *ListNode
	for _, num := range sumResult {
		res = &ListNode{Val: num, Next: res}
	}

	return res
}