package linklist

import (
	"fmt"
	"testing"
)

func TestBSTAfterOrderCheck(t *testing.T) {

	nums := []int{1,3,2,5,7,6,4}
	l := len(nums)
	ret := BSTAfterOrderCheck(nums,0, l-1)
	if ret == false {
		t.Errorf("BSTAfterOrderCheck error")
	}

	nums2 := []int{1,4,2,5,7,6,3}
	l2 := len(nums2)
	ret2 := BSTAfterOrderCheck(nums2,0, l2-1)
	if ret2 == true {
		t.Errorf("BSTAfterOrderCheck error")
	}

	nums3 := []int{4,5,3}
	l3 := len(nums3)
	ret3 := BSTAfterOrderCheck(nums3,0, l3-1)
	fmt.Println("ret3:", ret3)
	if ret3  != true {
		t.Errorf("BSTAfterOrderCheck error")
	}
}