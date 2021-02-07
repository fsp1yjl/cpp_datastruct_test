package sort

import (
	"math/rand"
	"testing"
)


func IntSliceCmp(a []int, b []int) bool {
	l1 := len(a)
	l2 := len(b)
	if l1 !=l2 {
		return false
	}
	for i:=0; i< l1;i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func SortedSliceCheck(n []int)  bool {
	l := len(n)
	for i:=0; i<l;i++ {
		if n[i] != i {
			return false
		}
	}
	return true
}
func TestQuickSort(t *testing.T) {

	numLength :=10
	for ; numLength< 5000; numLength++ {

		// rand.Perm会返回0-total数据的随机全排列
		permNums := rand.Perm(numLength+1)
		//fmt.Println("BEFORE:", permNums)

		QuickSort(permNums)
		//fmt.Println("AFTER:", permNums)
		if !SortedSliceCheck(permNums)  {
			t.Errorf("quick sort wrong with total:%d", numLength )
		}
	}
}