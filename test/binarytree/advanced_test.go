package binarytree

import (
	"github.com/uzmijnlm/go_advanced/main/binarytree"
	"testing"
)

func TestSerialize(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \   \
					3   4   8
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7, 8}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	result := binarytree.Serialize(node)
	expectedRes := "1_2_3_#_#_4_5_#_#_6_#_#_7_#_8_#_#_"
	if result != expectedRes {
		t.Fail()
	}
}

func TestDeSerialize(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \   \
					3   4   8
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7, 8}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	serialize := binarytree.Serialize(node)
	result := binarytree.DeSerialize(serialize)

	serialize2 := binarytree.Serialize(result)

	if serialize2 != serialize {
		t.Fail()
	}
}

func TestGetMaxBreadth(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \   \
					3   4   8
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7, 8}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	maxBreadth := binarytree.GetMaxBreadth(node)

	if maxBreadth != 3 {
		t.Fail()
	}

}

func TestFindLowestCommonAncestor(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \   \
					3   4   8
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7, 8}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	// 查看3和6的LCA
	node3 := node.Left.Left
	node6 := node.Left.Right.Right
	node2 := node.Left
	if binarytree.FindLowestCommonAncestor(node, node3, node6) != node2 {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, node3, node6) != node2 {
		t.Fail()
	}

	// 查看5和6的LCA
	node5 := node.Left.Right.Left
	node4 := node.Left.Right
	if binarytree.FindLowestCommonAncestor(node, node5, node6) != node4 {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, node5, node6) != node4 {
		t.Fail()
	}

	// 查看4和6的LCA
	if binarytree.FindLowestCommonAncestor(node, node4, node6) != node4 {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, node4, node6) != node4 {
		t.Fail()
	}

	// 查看3和8的LCA
	node8 := node.Right.Right
	if binarytree.FindLowestCommonAncestor(node, node3, node8) != node {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, node3, node8) != node {
		t.Fail()
	}

	// 查看相同节点的LCA
	if binarytree.FindLowestCommonAncestor(node, node3, node3) != node3 {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, node3, node3) != node3 {
		t.Fail()
	}

	// 查看不属于同一棵树的节点与3的LCA
	otherNode := binarytree.NewTreeNode(-1)
	if binarytree.FindLowestCommonAncestor(node, node3, otherNode) != nil {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, node3, otherNode) != nil {
		t.Fail()
	}

	// 其他边界情况
	if binarytree.FindLowestCommonAncestor(node, node3, nil) != nil {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, node3, nil) != nil {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestor(node, nil, nil) != nil {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(node, nil, nil) != nil {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestor(nil, node3, node4) != nil {
		t.Fail()
	}
	if binarytree.FindLowestCommonAncestorOptimized(nil, node3, node4) != nil {
		t.Fail()
	}

}

func TestGetSizeOfCompleteBinaryTree(t *testing.T) {
	var preOrder = make([]int, 0)
	var inOrder = make([]int, 0)
	var node *binarytree.TreeNode
	/*
					      6
					     / \
			            2   7
					  /  \  / \
					1     4 8  9
		           / \   / \
		          10  11 3   5
	*/
	preOrder = []int{6, 2, 1, 10, 11, 4, 3, 5, 7, 8, 9}
	inOrder = []int{10, 1, 11, 2, 3, 4, 5, 6, 8, 7, 9}
	node = binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	if binarytree.GetSizeOfCompleteBinaryTree(node) != 11 {
		t.Fail()
	}

	/*
					        6
					     / 	  \
			            2   	7
					  /  \  	/ \
					1     4 	8  9
		           / \   / \    /
		          10  11 3   5  12
	*/
	preOrder = []int{6, 2, 1, 10, 11, 4, 3, 5, 7, 8, 12, 9}
	inOrder = []int{10, 1, 11, 2, 3, 4, 5, 6, 12, 8, 7, 9}
	node = binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	if binarytree.GetSizeOfCompleteBinaryTree(node) != 12 {
		t.Fail()
	}

}

func TestGetFolds(t *testing.T) {
	var res []int
	var expectedRes []int
	res = binarytree.GetFolds(1)
	expectedRes = []int{0}
	if len(res) != len(expectedRes) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedRes[i] {
			t.Fail()
		}
	}

	res = binarytree.GetFolds(2)
	expectedRes = []int{0, 0, 1}
	if len(res) != len(expectedRes) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedRes[i] {
			t.Fail()
		}
	}

	res = binarytree.GetFolds(3)
	expectedRes = []int{0, 0, 1, 0, 0, 1, 1}
	if len(res) != len(expectedRes) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedRes[i] {
			t.Fail()
		}
	}
}

func TestGetSuccessor(t *testing.T) {
	/*
					        6
					     / 	  \
			            2   	7
					  /  \  	/ \
					1     4 	8  9
		           / \   / \    /
		          10  11 3   5  12
	*/
	preOrder := []int{6, 2, 1, 10, 11, 4, 3, 5, 7, 8, 12, 9}
	inOrder := []int{10, 1, 11, 2, 3, 4, 5, 6, 12, 8, 7, 9}
	node := BuildTreeGivenPreAndIn(preOrder, inOrder)

	// 10的后继节点
	node10 := node.Left.Left.Left
	node1 := node.Left.Left
	if binarytree.GetSuccessor(node10) != node1 {
		t.Fail()
	}

	// 11的后继节点
	node11 := node.Left.Left.Right
	node2 := node.Left
	if binarytree.GetSuccessor(node11) != node2 {
		t.Fail()
	}

	// 9的后继节点
	node9 := node.Right.Right
	if binarytree.GetSuccessor(node9) != nil {
		t.Fail()
	}

}

func BuildTreeGivenPreAndIn(preOrder []int, inOrder []int) *binarytree.SpecialTreeNode {
	return buildTreeGivenPreAndIn(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1)
}

func buildTreeGivenPreAndIn(preOrder []int, preL int, preR int, inOrder []int, inL int, inR int) *binarytree.SpecialTreeNode {
	if preL > preR || inL > inR {
		return nil
	}
	midValue := preOrder[preL]
	node := binarytree.NewSpecialTreeNode(midValue)
	midIndex := inL
	for inOrder[midIndex] != midValue {
		midIndex += 1
	}
	leftSize := midIndex - inL
	node.Left = buildTreeGivenPreAndIn(preOrder, preL+1, preL+leftSize, inOrder, inL, midIndex-1)
	if node.Left != nil {
		node.Left.Parent = node
	}
	node.Right = buildTreeGivenPreAndIn(preOrder, preL+leftSize+1, preR, inOrder, midIndex+1, inR)
	if node.Right != nil {
		node.Right.Parent = node
	}
	return node
}

func TestConvertToDoubleNode(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \   \
					3   4   8
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7, 8}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	start := binarytree.ConvertToDoubleNode(node)
	cur := start
	index := 0
	var end *binarytree.TreeNode
	for cur != nil {
		if cur.Value != inOrder[index] {
			t.Fail()
		}
		index += 1
		if cur.Right == nil {
			end = cur
		}
		cur = cur.Right
	}

	index = 7
	cur = end
	for cur != nil {
		if cur.Value != inOrder[index] {
			t.Fail()
		}
		index -= 1
		cur = cur.Left
	}

}

func TestPrintEdge(t *testing.T) {
	/*
							   1
						/			  \
					   2               3
		                \              / \
						 4             5  6
						/ \           / \
			           7   8         9   10
		                    \       /
		                    11      12
		                   /  \    /  \
		                  13  14  15   16
	*/
	preOrder := []int{1, 2, 4, 7, 8, 11, 13, 14, 3, 5, 9, 12, 15, 16, 10, 6}
	inOrder := []int{2, 7, 4, 8, 13, 11, 14, 1, 15, 12, 16, 9, 5, 10, 3, 6}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	res := binarytree.PrintEdge(node)
	expectedRes := []int{1, 2, 4, 7, 11, 13, 14, 15, 16, 12, 10, 6, 3}
	if len(res) != len(expectedRes) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedRes[i] {
			t.Fail()
		}
	}
}

func TestGetLongestPathWithSum(t *testing.T) {
	/*
						-3
			 		  /    \
					3		-9
		          /  \      / \
		         1    0    2   1
					 / \
		            1   6
	*/
	preOrder := []int{-3, 3, 1, 0, 1, 6, -9, 2, 1}
	inOrder := []int{1, 3, 1, 0, 6, -3, 2, -9, 1}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	if binarytree.GetMaxLenWithSum(node, 6) != 4 {
		t.Fail()
	}
	if binarytree.GetMaxLenWithSum(node, -9) != 1 {
		t.Fail()
	}
}

func TestGetMaxBST(t *testing.T) {
	/*
						 6
			 		  /      \
					1		   12
		          /  \        / \
		         0    3    10     13
						  /  \   / \
				        4    14 20  16
		              /  \  /  \
					2    5 11  15
	*/
	preOrder := []int{6, 1, 0, 3, 12, 10, 4, 2, 5, 14, 11, 15, 13, 20, 16}
	inOrder := []int{0, 1, 3, 6, 2, 4, 5, 10, 11, 14, 15, 12, 20, 13, 16}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	expectedPreOrder := []int{10, 4, 2, 5, 14, 11, 15}
	bst := binarytree.GetMaxBST(node)

	res := binarytree.PreOrderRecur(bst, make([]int, 0))

	if len(res) != len(expectedPreOrder) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedPreOrder[i] {
			t.Fail()
		}
	}
}

func TestGetMaxBSTopoSize(t *testing.T) {
	/*
						 6
			 		  /      \
					1		   12
		          /  \        / \
		         0    3    10     13
						  /  \   / \
				        4    14 20  16
		              /  \  /  \
					2    5 11  15
	*/
	preOrder := []int{6, 1, 0, 3, 12, 10, 4, 2, 5, 14, 11, 15, 13, 20, 16}
	inOrder := []int{0, 1, 3, 6, 2, 4, 5, 10, 11, 14, 15, 12, 20, 13, 16}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	res := binarytree.GetMaxBSTopoSize(node)
	if res != 8 {
		t.Fail()
	}
}

func TestTraverseZigZag(t *testing.T) {
	/*
							 1
				 		  /      \
						2		  3
			          /          / \
			         4          5    6
							  /  \
		                     7    8
	*/
	preOrder := []int{1, 2, 4, 3, 5, 7, 8, 6}
	inOrder := []int{4, 2, 1, 7, 5, 8, 3, 6}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	res := binarytree.TraverseZigZag(node)
	expectedRes := []int{1, 3, 2, 4, 5, 6, 8, 7}
	if len(res) != len(expectedRes) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedRes[i] {
			t.Fail()
		}
	}
}

func TestFindTwoErrorNodes(t *testing.T) {
	/*
					    7
					   / \
			          2   6
			         / \
					1   4
		               / \
		              3   5
	*/
	preOrder := []int{7, 2, 1, 4, 3, 5, 6}
	inOrder := []int{1, 2, 3, 4, 5, 7, 6}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	res := binarytree.FindTwoErrorNodes(node)
	if res[0].Value != 7 || res[1].Value != 6 {
		t.Fail()
	}

	/*
					    6
					   / \
			          5   7
			         / \
					1   4
		               / \
		              3   2
	*/
	preOrder = []int{6, 5, 1, 4, 3, 2, 7}
	inOrder = []int{1, 5, 3, 4, 2, 6, 7}
	node = binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	res = binarytree.FindTwoErrorNodes(node)
	if res[0].Value != 5 || res[1].Value != 2 {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	/*
						    1
		                 /      \
		                2        3                  2
		               / \      / \               /  \
		             4    5    6   7             4    5
		            / \   /                     /
		           8   9  10                   8
	*/
	preOrder := []int{1, 2, 4, 8, 9, 5, 10, 3, 6, 7}
	inOrder := []int{8, 4, 9, 2, 10, 5, 1, 6, 3, 7}
	node1 := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	preOrder = []int{2, 4, 8, 5}
	inOrder = []int{8, 4, 2, 5}
	node2 := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	if !binarytree.Contains(node1, node2) {
		t.Fail()
	}
}

func TestBuildBSTByPostOrder(t *testing.T) {
	/*
						      4
						   /     \
				          3       7
				         /      /
						1      5
		                 \     \
		                  2     6
	*/
	inOrder := []int{1, 2, 3, 4, 5, 6, 7}
	postOrder := []int{2, 1, 3, 6, 5, 7, 4}

	node := binarytree.BuildBSTByPostOrder(postOrder)
	expectedNode := binarytree.BuildTreeGivenInAndPost(inOrder, postOrder)

	preOrder := binarytree.PreOrderRecur(node, make([]int, 0))
	expectedPreOrder := binarytree.PreOrderRecur(expectedNode, make([]int, 0))
	if len(preOrder) != len(expectedPreOrder) {
		t.Fail()
	}
	for i := 0; i < len(preOrder); i++ {
		if preOrder[i] != expectedPreOrder[i] {
			t.Fail()
		}
	}
}

func TestBuildBSTBySortedArr(t *testing.T) {
	/*
				      4
				   /     \
		          2       6
		         / \     / \
				1   3   5   7
	*/
	inOrder := []int{1, 2, 3, 4, 5, 6, 7}
	postOrder := []int{1, 3, 2, 5, 7, 6, 4}

	node := binarytree.BuildBalancedBSTBySortedArr(inOrder)
	expectedNode := binarytree.BuildTreeGivenInAndPost(inOrder, postOrder)

	preOrder := binarytree.PreOrderRecur(node, make([]int, 0))
	expectedPreOrder := binarytree.PreOrderRecur(expectedNode, make([]int, 0))
	if len(preOrder) != len(expectedPreOrder) {
		t.Fail()
	}
	for i := 0; i < len(preOrder); i++ {
		if preOrder[i] != expectedPreOrder[i] {
			t.Fail()
		}
	}
}

func TestGetMaxDistance(t *testing.T) {
	/*
				      4
				   /     \
		          2       6
		         / \     / \
				1   3   5   7
	*/
	inOrder := []int{1, 2, 3, 4, 5, 6, 7}
	postOrder := []int{1, 3, 2, 5, 7, 6, 4}

	node := binarytree.BuildTreeGivenInAndPost(inOrder, postOrder)

	if binarytree.GetMaxDistance(node) != 5 {
		t.Fail()
	}
}
