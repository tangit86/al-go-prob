package riddles

import "fmt"

// Candies : Calculates the minimum amount of candies the teacher should hand out to kids
// n: input array size (kids count)
// arr: array with the grades of each kid
func Candies(n int, arr []int) int {
	result := make([]int, n)

	cur := 0

	for cur < n-1 {
		var curVal, curCandy = arr[cur], result[cur]
		var nextVal, nextCandy = arr[cur+1], result[cur+1]

		if nextVal > curVal && nextCandy <= curCandy {
			result[cur+1] = curCandy + 1
		}

		if nextVal < curVal && nextCandy >= curCandy {
			result[cur]++
			cur -= 2
			if cur < 0 {
				cur = 0
			}
		}

		cur++
	}

	k, cnt := 0, 0
	for k < n {
		cnt += result[k] + 1
		fmt.Printf("%v-", cnt)
		k++
	}
	fmt.Printf("\nResult: %v\n", cnt)
	return cnt
}
