package piscine

// BTreeApplyByLevel applies the function f to each node of the binary tree by level.
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// Apply the function f to the node's data
		_, _ = f(node.Data)

		// Enqueue the left and right children if they exist
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
