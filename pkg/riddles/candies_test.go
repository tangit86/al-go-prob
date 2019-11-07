package riddles

import "testing"

func TestCandies(t *testing.T) {
	res := Candies(3, []int{1, 2, 3})
	expectedCandies := 5
	if res != expectedCandies {
		t.Errorf("Test failed, Returnt %d , expected %d", res, expectedCandies)
	}
}
