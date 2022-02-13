package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBST(1, n)
}

func generateBST(start, end int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if start > end {
		trees = append(trees, nil)
		return trees
	}
	for i := start; i <= end; i++ {
		leftChildren, rightChildren := generateBST(start, i-1), generateBST(i+1, end)
		for _, leftChild := range leftChildren {
			for _, rightChild := range rightChildren {
				trees = append(trees, &TreeNode{Val: i, Left: leftChild, Right: rightChild})
			}
		}
	}
	return trees
}
