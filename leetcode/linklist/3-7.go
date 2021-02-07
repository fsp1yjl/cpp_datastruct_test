package linklist

/*
判断一个数组是否是二叉查找树后序遍历的结果

         4
       /  \
      2     6     =>    1-3-2-5-7-6-4
     / \   / \
    1   3  5  7
*/

func BSTAfterOrderCheck(arr []int, left int, right int) bool {
	if left == right {
		return true
	}

	middle := arr[right]

	// 找到第一个大于middle的节点
	var i int
	for i = left; i < right; i++ {
		if arr[i] > middle {
			break
		}
	}

	// 如果第一个大于middle之后的值有小于middle，则非法
	for j := i; j < right; j++ {
		if arr[j] < middle {
			return false
		}
	}

	if i == right || arr[i] < middle {
		//所有左侧都小于middle，即无右子树的情况
		return BSTAfterOrderCheck(arr, left, right-1)
	}
	if i == left {
		//第一个左侧即大于middle，表示无左子

		return BSTAfterOrderCheck(arr, left, right-1)
	}

	//左右都存在都情况
	return BSTAfterOrderCheck(arr, left, i-1) && BSTAfterOrderCheck(arr, i, right-1)

}
