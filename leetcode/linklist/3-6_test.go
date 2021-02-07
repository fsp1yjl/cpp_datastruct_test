package linklist

import (
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

func CreateTree()(root *Node){
	node4 := &Node{
		data:4,
	}

	node2 := &Node{
		data:2,
	}
	node6 := &Node{
		data:6,
	}

	node1 := &Node{
		data:1,
	}

	node3 := &Node{
		data:3,
	}
	node5 := &Node{
		data:5,
	}

	node7 := &Node{
		data:7,
	}

	node4.left = node2
	node4.right= node6
	node2.left = node1
	node2.right = node3
	node6.left = node5
	node6.right = node7

	return node4

}

func TestTree2DoubleList(t *testing.T) {
	// go test  -v -timeout 30s linklist -run Testtree2DoubleList
	root := CreateTree()

	//fmt.Println("rooot.data:", root.left.data)
	head, tail := Tree2DoubleList(root)

	if head == nil || tail == nil {
		t.Errorf("process error: head or tail is nil")
		return
	}
	if head.data != 1 || tail.data != 7 {
		t.Errorf("process error, wrong head or tail")
		return
	}

	nums1 := []int{1,2,3,4,5,6,7}
	nums2 := []int{7,6,5,4,3,2,1}

	n1 := make([]int,0)
	h := head
	for h!= nil {
		//fmt.Println(" h DATA:", h.data)
		n1 = append(n1,  h.data)
		h = h.right
	}
	if !IntSliceCmp(nums1, n1) {
		t.Errorf("slice left to right check error")
		return
	}

	tt := tail
	n2 := make([]int,0)

	for tt != nil {
		//fmt.Println("t DATA:", tt.data)
		n2 = append(n2,  tt.data)
		tt  = tt.left
	}

	if !IntSliceCmp(nums2, n2) {
		t.Errorf("slice right to left check error")
		return
	}

}

func TestTree2List(t *testing.T) {
	// go test  -v -timeout 30s linklist -run Testtree2DoubleList
	root := CreateTree()

	//fmt.Println("rooot.data:", root.left.data)
	head, tail := Tree2List(root)

	if head == nil || tail == nil {
		t.Errorf("process error: head or tail is nil")
		return
	}
	if head.data != 1 || tail.data != 7 {
		t.Errorf("process error, wrong head or tail")
		return
	}


	nums1 := []int{1,2,3,4,5,6,7}

	n1 := make([]int,0)
	h := head
	for h!= nil {
		//fmt.Println(" h DATA:", h.data)
		n1 = append(n1,  h.data)
		h = h.right
	}
	if !IntSliceCmp(nums1, n1) {
		t.Errorf("slice left to right check error")
		return
	}

}