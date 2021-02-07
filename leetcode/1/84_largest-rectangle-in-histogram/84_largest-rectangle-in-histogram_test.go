package l84

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println("h")
	nums1 := []int{2, 1, 4, 5, 1, 3, 3}
	ret1 := largestRectangleArea(nums1)
	if ret1 != 8 {
		t.Error("error nums1")
	}

	nums2 := []int{2, 1, 5, 6, 2, 3}
	ret2 := largestRectangleArea(nums2)
	if ret2 != 10 {
		t.Error("error nums2")
	}
}
