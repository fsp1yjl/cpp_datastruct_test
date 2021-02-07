package linklist

type Node struct {
	data  int
	left  *Node
	right *Node
}
/*

         4
       /  \
      2     6     =>    1-2-3-4-5-6-7
     / \   / \
    1   3  5  7

*/

// 2叉数按中序遍历的顺序放入双向链表
func Tree2DoubleList(root *Node) (head *Node, tail *Node) {

	//fmt.Println("root.value:", root.data)
	if root == nil {
		return nil, nil
	}

	if root.left == nil && root.right == nil {
		n :=&Node{
			data: root.data,
		}
		return n, n
	}

	// 无左子的情况
	if root.left == nil {
		head := &Node{
			data:root.data,
			left: nil,
			right:nil,
		}
		rHead,rTail := Tree2DoubleList(root.right)

		rHead.left = head
		head.right = rHead
		return head, rTail
	}


	//无右子的情况
	if root.right == nil {
		tail := &Node{
			data:root.data,
			left: nil,
			right:nil,
		}
		lHead,lTail := Tree2DoubleList(root.left)
		lTail.right = tail
		tail.left = lTail
		return lHead,lTail
	}

	//左右均有的情况
    cur := &Node{
    	data: root.data,
	}
	rHead,rTail := Tree2DoubleList(root.right)
	lHead,lTail := Tree2DoubleList(root.left)
	cur.left = lTail
	cur.right = rHead
	rHead.left = cur
	lTail.right = cur
	return lHead, rTail



	return nil, nil
}

// 2叉数按中序遍历的顺序放入单向链表
func Tree2List(root *Node) (head *Node, tail *Node) {

	//fmt.Println("root.value:", root.data)
	if root == nil {
		return nil, nil
	}

	if root.left == nil && root.right == nil {
		n :=&Node{
			data: root.data,
		}
		return n, n
	}

	// 左子树空
	if root.left == nil {
		head := &Node{
			data:root.data,
			right:nil,
		}
		rHead,rTail := Tree2List(root.right)
		head.right = rHead
		return head, rTail
	}

	// 右空
	if root.right == nil {
		tail := &Node{
			data:root.data,
			right:nil,
		}
		lHead,lTail := Tree2DoubleList(root.left)
		lTail.right = tail
		return lHead,lTail
	}

	cur := &Node{
		data: root.data,
	}
	rHead,rTail := Tree2DoubleList(root.right)
	lHead,lTail := Tree2DoubleList(root.left)

	cur.right = rHead
	lTail.right = cur
	return lHead, rTail



	return nil, nil
}

