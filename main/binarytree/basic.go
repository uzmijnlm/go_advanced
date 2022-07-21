package binarytree

import (
	"github.com/uzmijnlm/go_advanced/main/basic"
	"math"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Value: val}
}

func PreOrderRecur(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = append(result, node.Value)
	result = PreOrderRecur(node.Left, result)
	result = PreOrderRecur(node.Right, result)
	return result
}

func PreOrderIter(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	stack := basic.NewStack()
	stack.Push(node)
	for !stack.IsEmpty() {
		current := stack.Pop().(*TreeNode)
		result = append(result, current.Value)
		if current.Right != nil {
			stack.Push(current.Right)
		}
		if current.Left != nil {
			stack.Push(current.Left)
		}
	}

	return result
}

func InOrderRecur(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = InOrderRecur(node.Left, result)
	result = append(result, node.Value)
	result = InOrderRecur(node.Right, result)
	return result
}

func InOrderIter(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	stack := basic.NewStack()
	stack.Push(node)
	cur := node.Left
	for cur != nil || !stack.IsEmpty() {
		// 先将子树所有左节点都放入stack
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}

		// 从子树最左节点开始处理，这个节点没有左节点
		cur = stack.Pop().(*TreeNode)
		result = append(result, cur.Value)

		// 开始处理右子树
		cur = cur.Right
	}
	return result
}

func PostOrderRecur(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = PostOrderRecur(node.Left, result)
	result = PostOrderRecur(node.Right, result)
	result = append(result, node.Value)
	return result
}

func PostOrderIter(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	stack1 := basic.NewStack()
	stack2 := basic.NewStack()

	stack1.Push(node)
	for !stack1.IsEmpty() {
		cur := stack1.Pop().(*TreeNode)
		if cur.Left != nil {
			stack1.Push(cur.Left)
		}
		if cur.Right != nil {
			stack1.Push(cur.Right)
		}
		stack2.Push(cur)
	}
	for !stack2.IsEmpty() {
		cur := stack2.Pop().(*TreeNode)
		result = append(result, cur.Value)
	}
	return result
}

// BuildTreeGivenPreAndIn 注意：无重复元素
func BuildTreeGivenPreAndIn(preOrder []int, inOrder []int) *TreeNode {
	return buildTreeGivenPreAndIn(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1)
}

// 思路：
// 前序遍历 1  (2 4 5 8 9)  (3 6 10 7)
// 中序遍历 (4 2 8 5 9)  1  (6 10 3 7)
// 根据前序遍历找到中间节点，然后在中序遍历中找到这个节点，以此计算出左右子树的大小
func buildTreeGivenPreAndIn(preOrder []int, preL int, preR int, inOrder []int, inL int, inR int) *TreeNode {
	if preL > preR || inL > inR {
		return nil
	}
	midValue := preOrder[preL]
	node := NewTreeNode(midValue)
	midIndex := inL
	for inOrder[midIndex] != midValue {
		midIndex += 1
	}
	leftSize := midIndex - inL
	node.Left = buildTreeGivenPreAndIn(preOrder, preL+1, preL+leftSize, inOrder, inL, midIndex-1)
	node.Right = buildTreeGivenPreAndIn(preOrder, preL+leftSize+1, preR, inOrder, midIndex+1, inR)
	return node
}

// BuildTreeGivenInAndPost 注意：无重复元素
func BuildTreeGivenInAndPost(inOrder []int, postOrder []int) *TreeNode {
	return buildTreeGivenInAndPost(inOrder, 0, len(inOrder)-1, postOrder, 0, len(postOrder)-1)
}

// 思路与给定前序遍历和中序遍历一致
func buildTreeGivenInAndPost(inOrder []int, inL int, inR int, postOrder []int, postL int, postR int) *TreeNode {
	if inL > inR || postL > postR {
		return nil
	}

	midValue := postOrder[postR]
	node := NewTreeNode(midValue)
	midIndex := inL
	for inOrder[midIndex] != midValue {
		midIndex += 1
	}
	leftSize := midIndex - inL
	rightSize := inR - midIndex

	node.Left = buildTreeGivenInAndPost(inOrder, inL, midIndex-1, postOrder, postL, postL+leftSize-1)
	node.Right = buildTreeGivenInAndPost(inOrder, midIndex+1, inR, postOrder, postL+leftSize, postL+leftSize+rightSize-1)
	return node
}

// BuildTreeGivenPreAndPost 无重复元素
func BuildTreeGivenPreAndPost(preOrder []int, postOrder []int) *TreeNode {
	return buildTreeGivenPreAndPost(0, 0, len(preOrder), preOrder, postOrder)
}

// 思路：
// 前序遍历为：（根节点）（前序遍历左分支）（前序遍历右分支）
// 后序遍历为：（后序遍历左分支）（后序遍历右分支）（根节点）
// 只要有办法知道左分支的节点个数，就可以递归地进行构造
// 用preOrder[i0:i0+n]表示preOrder中的左分支，用postOrder[i1:i1+n]表示postOrder中的左分支
// i0表示preOrder左分支的开头，i1表示postOrder左分支的开头，n表示左分支的长度
// 关键要记住的是postOrder左分支最后一个节点等于preOrder左分支第一个节点
func buildTreeGivenPreAndPost(i0 int, i1 int, n int, preOrder []int, postOrder []int) *TreeNode {
	if n == 0 {
		return nil
	}
	node := NewTreeNode(preOrder[i0])
	if n == 1 {
		return node
	}

	l := 1 // 左分支的长度
	preOrderFirstLeft := preOrder[i0+1]
	for l < n {
		if postOrder[i1+l-1] == preOrderFirstLeft { // postOrder左分支最后一个节点等于preOrder左分支第一个节点
			break
		}
		l += 1
	}

	node.Left = buildTreeGivenPreAndPost(i0+1, i1, l, preOrder, postOrder)
	node.Right = buildTreeGivenPreAndPost(i0+l+1, i1+l, n-1-l, preOrder, postOrder)
	return node
}

func TraverseBreadthFirst(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	res := make([]int, 0)
	queue := basic.NewQueue()
	queue.Push(node)
	for !queue.IsEmpty() {
		cur := queue.Pop().(*TreeNode)
		res = append(res, cur.Value)
		if cur.Left != nil {
			queue.Push(cur.Left)
		}
		if cur.Right != nil {
			queue.Push(cur.Right)
		}
	}

	return res
}

func TraverseDepthFirst(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	stack := basic.NewStack()
	stack.Push(node)
	res := make([]int, 0)
	for !stack.IsEmpty() {
		cur := stack.Pop().(*TreeNode)
		res = append(res, cur.Value)
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}
	return res
}

// IsSearchBinaryTree 判断是否为平衡二叉树。对于任意一个节点，左子树一定比它小，右子树一定比它大
func IsSearchBinaryTree(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if node.Left != nil {
		if node.Left.Value >= node.Value {
			return false
		}
	}
	if node.Right != nil {
		if node.Right.Value <= node.Value {
			return false
		}
	}

	return IsSearchBinaryTree(node.Left) && IsSearchBinaryTree(node.Right)
}

// IsCompleteBinaryTree 判断是否为完全二叉树。要么全满，要么从左到右依次填满
// 宽度优先遍历。在遍历过程中判断：
// a.如果一个节点只有右节点没有左节点，则一定不是满二叉树
// b.如果一个节点只有左节点，那么后面所有节点都不能有子节点
func IsCompleteBinaryTree(node *TreeNode) bool {
	if node == nil {
		return true
	}

	queue := basic.NewQueue()
	queue.Push(node)
	meetLastLeft := false
	for !queue.IsEmpty() {
		cur := queue.Pop().(*TreeNode)
		if cur.Left == nil && cur.Right != nil {
			return false
		}
		if meetLastLeft {
			if cur.Left != nil || cur.Right != nil {
				return false
			}
		} else {
			if (cur.Left != nil && cur.Right == nil) || (cur.Left == nil && cur.Right == nil) {
				meetLastLeft = true
			}
		}
		if cur.Left != nil {
			queue.Push(cur.Left)
		}
		if cur.Right != nil {
			queue.Push(cur.Right)
		}
	}
	return true
}

// IsBalancedBinaryTree 判断是否为平衡二叉树。定义：左右子树都是平衡二叉树，且高度差不超过1
func IsBalancedBinaryTree(node *TreeNode) bool {
	isBalanced, _ := isBalancedBinaryTree(node)
	return isBalanced
}

func isBalancedBinaryTree(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := isBalancedBinaryTree(node.Left)
	isRightBalanced, rightHeight := isBalancedBinaryTree(node.Right)
	diff := math.Abs(float64(leftHeight - rightHeight))
	height := int(math.Max(float64(leftHeight), float64(rightHeight)))
	return isLeftBalanced && isRightBalanced && (diff <= 1), height + 1
}
