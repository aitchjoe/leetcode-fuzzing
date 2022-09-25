package solution

func longestSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	arr := []int{}
	repeat := 0
	prev := -1
	for _, n := range nums {
		if n != prev && prev != -1 {
			arr = mergeNums(arr, prev, repeat)
			repeat = 1
		} else {
			repeat++
		}
		prev = n
	}

	arr = mergeNums(arr, prev, repeat)
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		if nums[0] == 0 || nums[len(nums)-1] == 0 { // [0,0,0] or [1,1,1,0]
			return arr[0]
		} else { // [1,1,1]
			return arr[0] - 1
		}
	}

	max := 0
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i] + arr[i+1]
		if sum > max {
			max = sum
		}
	}
	return max
}

func mergeNums(arr []int, num, repeat int) []int {
	if num == 1 {
		arr = append(arr, repeat)
	} else if num == 0 && repeat > 1 {
		arr = append(arr, 0)
	}
	return arr
}
