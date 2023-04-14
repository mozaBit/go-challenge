package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lc := BTreeLevelCount(root.Left)
	rc := BTreeLevelCount(root.Right)
	if lc > rc {
		return lc + 1
	}
	return rc + 1
}
