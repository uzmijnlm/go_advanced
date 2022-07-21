package binarytree

import (
	"fmt"
	"github.com/uzmijnlm/go_advanced/main/binarytree"
	"testing"
)

func TestPreOrder(t *testing.T) {
	var node *binarytree.TreeNode
	var res1 []int
	var res2 []int
	res1 = binarytree.PreOrderRecur(node, res1)
	res2 = binarytree.PreOrderIter(node, res2)
	fmt.Println()
	if res1 != nil || res2 != nil {
		t.Fail()
	}

	node = binarytree.NewTreeNode(1)
	res1 = binarytree.PreOrderRecur(node, res1)
	res2 = binarytree.PreOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 1 || res1[0] != 1 || len(res2) != 1 || res2[0] != 1 {
		t.Fail()
	}

	/*
				    1
				   /
		          2
		         /
				3
	*/
	node.Left = binarytree.NewTreeNode(2)
	node.Left.Left = binarytree.NewTreeNode(3)
	res1 = make([]int, 0)
	res2 = make([]int, 0)
	res1 = binarytree.PreOrderRecur(node, res1)
	res2 = binarytree.PreOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 3 || res1[0] != 1 || res1[1] != 2 || res1[2] != 3 {
		t.Fail()
	}
	if len(res2) != 3 || res2[0] != 1 || res2[1] != 2 || res2[2] != 3 {
		t.Fail()
	}

	/*
				    1
				   / \
		          2   5
		         / \
				3   4
	*/
	node.Right = binarytree.NewTreeNode(5)
	node.Left.Right = binarytree.NewTreeNode(4)
	res1 = make([]int, 0)
	res2 = make([]int, 0)
	res1 = binarytree.PreOrderRecur(node, res1)
	res2 = binarytree.PreOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 5 || len(res2) != 5 {
		t.Fail()
	}
	for i := 0; i < 5; i++ {
		if res1[i] != i+1 || res2[i] != i+1 {
			t.Fail()
		}
	}
}

func TestInOrder(t *testing.T) {
	var node *binarytree.TreeNode
	var res1 []int
	var res2 []int
	res1 = binarytree.InOrderRecur(node, res1)
	res2 = binarytree.InOrderIter(node, res2)
	fmt.Println()
	if res1 != nil || res2 != nil {
		t.Fail()
	}

	node = binarytree.NewTreeNode(1)
	res1 = binarytree.InOrderRecur(node, res1)
	res2 = binarytree.InOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 1 || res1[0] != 1 {
		t.Fail()
	}
	if len(res2) != 1 || res2[0] != 1 {
		t.Fail()
	}

	/*
				    1
				   /
		          2
		         /
				3
	*/
	node.Left = binarytree.NewTreeNode(2)
	node.Left.Left = binarytree.NewTreeNode(3)
	res1 = make([]int, 0)
	res2 = make([]int, 0)
	res1 = binarytree.InOrderRecur(node, res1)
	res2 = binarytree.InOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 3 || len(res2) != 3 {
		t.Fail()
	}
	expectedRes := []int{3, 2, 1}
	for i := 0; i < len(expectedRes); i++ {
		if res1[i] != expectedRes[i] || res2[i] != expectedRes[i] {
			t.Fail()
		}
	}

	/*
				    1
				   / \
		          2   5
		         / \
				3   4
	*/
	node.Right = binarytree.NewTreeNode(5)
	node.Left.Right = binarytree.NewTreeNode(4)
	res1 = make([]int, 0)
	res2 = make([]int, 0)
	res1 = binarytree.InOrderRecur(node, res1)
	res2 = binarytree.InOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 5 || len(res2) != 5 {
		t.Fail()
	}
	expectedRes = []int{3, 2, 4, 1, 5}
	for i := 0; i < len(expectedRes); i++ {
		if res1[i] != expectedRes[i] || res2[i] != expectedRes[i] {
			t.Fail()
		}
	}
}

func TestPostOrder(t *testing.T) {
	var node *binarytree.TreeNode
	var res1 []int
	var res2 []int
	res1 = binarytree.PostOrderRecur(node, res1)
	res2 = binarytree.PostOrderIter(node, res2)
	fmt.Println()
	if res1 != nil || res2 != nil {
		t.Fail()
	}

	node = binarytree.NewTreeNode(1)
	res1 = binarytree.PostOrderRecur(node, res1)
	res2 = binarytree.PostOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 1 || res1[0] != 1 {
		t.Fail()
	}
	if len(res2) != 1 || res2[0] != 1 {
		t.Fail()
	}

	/*
				    1
				   /
		          2
		         /
				3
	*/
	node.Left = binarytree.NewTreeNode(2)
	node.Left.Left = binarytree.NewTreeNode(3)
	res1 = make([]int, 0)
	res2 = make([]int, 0)
	res1 = binarytree.PostOrderRecur(node, res1)
	res2 = binarytree.PostOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 3 || len(res2) != 3 {
		t.Fail()
	}
	expectedRes := []int{3, 2, 1}
	for i := 0; i < len(expectedRes); i++ {
		if res1[i] != expectedRes[i] || res2[i] != expectedRes[i] {
			t.Fail()
		}
	}

	/*
				    1
				   / \
		          2   5
		         / \
				3   4
	*/
	node.Right = binarytree.NewTreeNode(5)
	node.Left.Right = binarytree.NewTreeNode(4)
	res1 = make([]int, 0)
	res2 = make([]int, 0)
	res1 = binarytree.PostOrderRecur(node, res1)
	res2 = binarytree.PostOrderIter(node, res2)
	fmt.Println()
	if len(res1) != 5 || len(res2) != 5 {
		t.Fail()
	}
	expectedRes = []int{3, 4, 2, 5, 1}
	for i := 0; i < len(expectedRes); i++ {
		if res1[i] != expectedRes[i] || res2[i] != expectedRes[i] {
			t.Fail()
		}
	}
}

func TestBuildTreeGivenPreAndIn(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \
					3   4
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7}
	postOrder := []int{3, 5, 6, 4, 2, 7, 1}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)
	result := make([]int, 0)
	result = binarytree.PreOrderRecur(node, result)
	if len(result) != len(preOrder) {
		t.Fail()
	}
	for i := 0; i < len(preOrder); i++ {
		if result[i] != preOrder[i] {
			t.Fail()
		}
	}

	result = make([]int, 0)
	result = binarytree.InOrderRecur(node, result)
	if len(result) != len(inOrder) {
		t.Fail()
	}
	for i := 0; i < len(inOrder); i++ {
		if result[i] != inOrder[i] {
			t.Fail()
		}
	}

	result = make([]int, 0)
	result = binarytree.PostOrderIter(node, result)
	if len(result) != len(postOrder) {
		t.Fail()
	}
	for i := 0; i < len(postOrder); i++ {
		if result[i] != postOrder[i] {
			t.Fail()
		}
	}
}

func TestBuildTreeGivenInAndPost(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \
					3   4
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7}
	postOrder := []int{3, 5, 6, 4, 2, 7, 1}
	node := binarytree.BuildTreeGivenInAndPost(inOrder, postOrder)
	result := make([]int, 0)
	result = binarytree.InOrderRecur(node, result)
	if len(result) != len(inOrder) {
		t.Fail()
	}
	for i := 0; i < len(inOrder); i++ {
		if result[i] != inOrder[i] {
			t.Fail()
		}
	}

	result = make([]int, 0)
	result = binarytree.PostOrderIter(node, result)
	if len(result) != len(postOrder) {
		t.Fail()
	}
	for i := 0; i < len(postOrder); i++ {
		if result[i] != postOrder[i] {
			t.Fail()
		}
	}

	result = make([]int, 0)
	result = binarytree.PreOrderIter(node, result)
	if len(result) != len(preOrder) {
		t.Fail()
	}
	for i := 0; i < len(preOrder); i++ {
		if result[i] != preOrder[i] {
			t.Fail()
		}
	}
}

func TestBuildTreeGivenPreAndPost(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \
					3   4
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7}
	postOrder := []int{3, 5, 6, 4, 2, 7, 1}
	node := binarytree.BuildTreeGivenPreAndPost(preOrder, postOrder)
	result := make([]int, 0)
	result = binarytree.InOrderRecur(node, result)
	if len(result) != len(inOrder) {
		t.Fail()
	}
	for i := 0; i < len(inOrder); i++ {
		if result[i] != inOrder[i] {
			t.Fail()
		}
	}

	result = make([]int, 0)
	result = binarytree.PostOrderIter(node, result)
	if len(result) != len(postOrder) {
		t.Fail()
	}
	for i := 0; i < len(postOrder); i++ {
		if result[i] != postOrder[i] {
			t.Fail()
		}
	}

	result = make([]int, 0)
	result = binarytree.PreOrderIter(node, result)
	if len(result) != len(preOrder) {
		t.Fail()
	}
	for i := 0; i < len(preOrder); i++ {
		if result[i] != preOrder[i] {
			t.Fail()
		}
	}
}

func TestTraverseBreadthFirst(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \
					3   4
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	res := binarytree.TraverseBreadthFirst(node)
	expectedRes := []int{1, 2, 7, 3, 4, 5, 6}
	if len(res) != len(expectedRes) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedRes[i] {
			t.Fail()
		}
	}

}

func TestTraverseDepthFirst(t *testing.T) {
	/*
					    1
					   / \
			          2   7
			         / \
					3   4
		               / \
		              5   6
	*/
	preOrder := []int{1, 2, 3, 4, 5, 6, 7}
	inOrder := []int{3, 2, 5, 4, 6, 1, 7}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	res := binarytree.TraverseDepthFirst(node)
	expectedRes := []int{1, 2, 3, 4, 5, 6, 7}
	if len(res) != len(expectedRes) {
		t.Fail()
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expectedRes[i] {
			t.Fail()
		}
	}

}

func TestIsSearchBinaryTree(t *testing.T) {
	/*
					    6
					   / \
			          2   7
			         / \
					1   4
		               / \
		              3   5
	*/
	preOrder := []int{6, 2, 1, 4, 3, 5, 7}
	inOrder := []int{1, 2, 3, 4, 5, 6, 7}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)
	if !binarytree.IsSearchBinaryTree(node) {
		t.Fail()
	}

	/*
					    1
					   / \
			          2   7
			         / \
					3   4
		               / \
		              5   6
	*/
	preOrder = []int{1, 2, 3, 4, 5, 6, 7}
	inOrder = []int{3, 2, 5, 4, 6, 1, 7}
	node = binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)
	if binarytree.IsSearchBinaryTree(node) {
		t.Fail()
	}

}

func TestIsFullBinaryTree(t *testing.T) {
	/*
					    6
					   / \
			          2   7
			         / \
					1   4
		               / \
		              3   5
	*/
	preOrder := []int{6, 2, 1, 4, 3, 5, 7}
	inOrder := []int{1, 2, 3, 4, 5, 6, 7}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	if binarytree.IsCompleteBinaryTree(node) == true {
		t.Fail()
	}

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

	if binarytree.IsCompleteBinaryTree(node) == false {
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

	if binarytree.IsCompleteBinaryTree(node) == false {
		t.Fail()
	}

}

func TestIsBalancedBinaryTree(t *testing.T) {
	/*
					    6
					   / \
			          2   7
			         / \
					1   4
		               / \
		              3   5
	*/
	preOrder := []int{6, 2, 1, 4, 3, 5, 7}
	inOrder := []int{1, 2, 3, 4, 5, 6, 7}
	node := binarytree.BuildTreeGivenPreAndIn(preOrder, inOrder)

	if binarytree.IsBalancedBinaryTree(node) == true {
		t.Fail()
	}

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

	if binarytree.IsBalancedBinaryTree(node) == false {
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

	if binarytree.IsBalancedBinaryTree(node) == false {
		t.Fail()
	}

}
