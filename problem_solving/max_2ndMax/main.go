package main

func getMaxNsecondMax(nums []int) [2]int {
	var max int = 0
	var secondMax int = 0
	if nums[0] > nums[1] {
		max = nums[0]
		secondMax = nums[1]
	} else {
		max = nums[1]
		secondMax = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > secondMax {
			if nums[i] <= max {
				secondMax = nums[i]
			} else {
				secondMax, max = max, nums[i]
			}
		}
	}

	return
}

// func getMaxNsecondMax(nums []int) (max int, secondMax int) {

// 	if nums[0] > nums[1] {
// 		max = nums[0]
// 		secondMax = nums[1]
// 	} else {
// 		max = nums[1]
// 		secondMax = nums[0]
// 	}
// 	for i := 2; i < len(nums); i++ {
// 		if nums[i] > secondMax {
// 			if nums[i] <= max {
// 				secondMax = nums[i]
// 			} else {
// 				secondMax, max = max, nums[i]
// 			}
// 		}
// 	}

// 	return max, secondMax
// }

func main() {
	nums := []int{1, 5, 87, 45, 8, 8}
	// max := nums[0] // assume first value is the smallest

	// for _, value := range nums {
	// 	if value > max {
	// 		max = value // found another smaller value, replace previous value in max
	// 	}
	// }
	getMaxNsecondMax(nums)
	//max, secondMax := getMaxNsecondMax(nums)
	//fmt.Println(max, secondMax)
	//getMax := getMaxNsecondMax(nums)

	//fmt.Println("Tow Oldest ages are: ", getMax)
}
