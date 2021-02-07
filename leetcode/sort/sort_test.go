package sort

import (
	"math/rand"
	"testing"
)




func TestSelectSort(t *testing.T) {

	numLength :=10
	for ; numLength< 5000; numLength++ {

		// rand.Perm会返回0-total数据的随机全排列
		permNums := rand.Perm(numLength+1)
		//fmt.Println("BEFORE1:", permNums)

		nums := SelectSort(permNums)
		//fmt.Println("AFTER:", nums)
		if !SortedSliceCheck(nums)  {
			t.Errorf("quick sort wrong with total:%d", numLength )
		}
	}
}

func TestBubbleSort(t *testing.T) {

	numLength :=10
	for ; numLength< 5000; numLength++ {

		// rand.Perm会返回0-total数据的随机全排列
		permNums := rand.Perm(numLength+1)
		//fmt.Println("BEFORE1:", permNums)

		nums := BubbleSort(permNums)
		//fmt.Println("AFTER:", nums)
		if !SortedSliceCheck(nums)  {
			t.Errorf("quick sort wrong with total:%d", numLength )
		}
	}
}

func TestInsertSort(t *testing.T) {

	numLength :=10
	for ; numLength< 5000; numLength++ {

		// rand.Perm会返回0-total数据的随机全排列
		permNums := rand.Perm(numLength+1)
		//fmt.Println("BEFORE1:", permNums)

		nums := InsertSort(permNums)
		//fmt.Println("AFTER:", nums)
		if !SortedSliceCheck(nums)  {
			t.Errorf("quick sort wrong with total:%d", numLength )
		}
	}
}