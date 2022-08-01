package binarytree

import (
	"github.com/uzmijnlm/go_advanced/main/basic"
	"math"
	"strconv"
	"strings"
)

// Serialize 用前序遍历进行序列化
func Serialize(node *TreeNode) string {
	if node == nil {
		return "#_"
	}

	str := strconv.Itoa(node.Value) + "_"
	str = str + Serialize(node.Left)
	str = str + Serialize(node.Right)
	return str
}

// DeSerialize 用前序遍历进行反序列化
func DeSerialize(str string) *TreeNode {
	splits := strings.Split(str, "_")
	queue := basic.NewQueue()
	for i := 0; i < len(splits); i++ {
		queue.Push(splits[i])
	}

	return deSerialize(queue)

}

func deSerialize(queue *basic.Queue) *TreeNode {
	if queue.IsEmpty() {
		return nil
	}
	ele := queue.Pop()
	if ele == "#" {
		return nil
	}
	num, err := strconv.Atoi(ele.(string))
	if err != nil {
		return nil
	}
	node := NewTreeNode(num)
	node.Left = deSerialize(queue)
	node.Right = deSerialize(queue)
	return node
}

// GetMaxBreadth 核心思路是宽度有限遍历，并要知道何时到了下一层
func GetMaxBreadth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	var curEnd = node     // 当前层最后一个节点，初始化为头节点
	var nextEnd *TreeNode // 下一层最后一个节点，初始为nil
	var curLevelBreadth int
	var maxBreadth int

	queue := basic.NewQueue()
	queue.Push(node)
	for !queue.IsEmpty() {
		cur := queue.Pop().(*TreeNode)
		if cur.Left != nil {
			queue.Push(cur.Left)
			nextEnd = cur.Left
		}
		if cur.Right != nil {
			queue.Push(cur.Right)
			nextEnd = cur.Right
		}
		curLevelBreadth += 1
		if cur == curEnd { // 弹出的节点为当前层最后一个节点，则更新curLevelBreadth、curEnd和nextEnd
			maxBreadth = int(math.Max(float64(maxBreadth), float64(curLevelBreadth)))
			curLevelBreadth = 0
			curEnd = nextEnd
			nextEnd = nil
		}
	}
	return maxBreadth
}

// FindLowestCommonAncestor
// 思路：
// 1.先遍历，遍历时维护每个节点与父节点的关系，放入map中
// 2.构造一个set，根据map将node1的所有祖先节点（包括自己）放入到set中
// 3.自下而上判断node2及其祖先节点是否在set中，返回第一个在set中的节点就是LCA
func FindLowestCommonAncestor(head *TreeNode, node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if head == nil {
		return nil
	}
	son2Father := make(map[*TreeNode]*TreeNode)
	stack := basic.NewStack()

	stack.Push(head)
	for !stack.IsEmpty() {
		cur := stack.Pop().(*TreeNode)
		if cur.Left != nil {
			stack.Push(cur.Left)
			son2Father[cur.Left] = cur
		}
		if cur.Right != nil {
			stack.Push(cur.Right)
			son2Father[cur.Right] = cur
		}
	}

	ancestorForNode1 := make(map[*TreeNode]interface{})
	son1 := node1
	for {
		ancestorForNode1[son1] = nil
		father1, ok := son2Father[son1]
		if !ok {
			break
		}
		son1 = father1
	}

	son2 := node2
	for {
		_, ok := ancestorForNode1[son2] // 第二个节点及祖先节点是否在集合中
		if ok {                         // 存在于集合中。这个就是最低的公共祖先节点
			return son2
		} else { // 不存在。继续向上找第二个节点的祖先节点
			son2, ok = son2Father[son2]
			if !ok { // 找到头节点也没找到，则这个节点与前一个节点不在同一棵树
				return nil
			}
		}
	}
}

func FindLowestCommonAncestorOptimized(head *TreeNode, node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if head == nil {
		return nil
	}

	if !inSameTree(head, node1, node2) {
		return nil
	}

	return findLCA(head, node1, node2)

}

func findLCA(head *TreeNode, node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if head == nil || head == node1 || head == node2 {
		return head
	}

	left := findLCA(head.Left, node1, node2)
	right := findLCA(head.Right, node1, node2)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return head
}

func inSameTree(head *TreeNode, node1 *TreeNode, node2 *TreeNode) bool {
	meet1 := false
	meet2 := false
	stack := basic.NewStack()
	stack.Push(head)
	for !stack.IsEmpty() {
		cur := stack.Pop().(*TreeNode)
		if cur == node1 {
			meet1 = true
		}
		if cur == node2 {
			meet2 = true
		}
		if meet1 && meet2 {
			break
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
	}
	return meet1 && meet2
}

// GetSizeOfCompleteBinaryTree 时间复杂度低于O(N)
func GetSizeOfCompleteBinaryTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	depth := getDeepestLayer(1, node)
	return getSizeOfCompleteBinaryTree(node, 1, depth)
}

// 核心思路：
// a.如果一个节点的右子树的最左节点到达了以这个节点为头节点的树的最大深度，则这个节点的左子树必是满二叉树
// b.如果一个节点的右子树的最左节点没有到达以这个节点为头节点的树的最大深度，则这个节点的右子树必是满二叉树
// 定义下面的函数，表示求以node为头节点的树的节点树，startLayer表示node所在层数，depth为整棵树最大深度
func getSizeOfCompleteBinaryTree(node *TreeNode, startLayer int, depth int) int {
	if startLayer == depth {
		return 1
	}

	if getDeepestLayer(startLayer+1, node.Right) == depth { // node节点的右子树的最左节点能够到达最后一层
		leftSize := 1<<(depth-startLayer) - 1                                     // 左子树是满二叉树，可以直接用公式计算
		rightSize := getSizeOfCompleteBinaryTree(node.Right, startLayer+1, depth) // 右子树递归计算
		return leftSize + rightSize + 1
	} else { // node节点的右子树的最左节点不能到达最后一层
		leftSize := getSizeOfCompleteBinaryTree(node.Left, startLayer+1, depth) // 左子树递归计算
		rightSize := 1<<(depth-1-startLayer) - 1                                // 右子树的最大深度一定为depth-1
		return leftSize + rightSize + 1
	}
}

func getDeepestLayer(startLayer int, node *TreeNode) int {
	layer := startLayer
	cur := node
	for cur != nil {
		layer += 1
		cur = cur.Left
	}
	return layer - 1
}

// GetFolds 一张纸条对折n次，返回从左到右的折痕。0表示凹，1表示凸
// 第一次对折后，为					0
// 第二次对折后，第二次的折痕为		0		1
// 第三次对折后，第三次的折痕为   0    1   0   1
// 按照对折的顺序，从上到下可以形成一棵树。每次对折后中序遍历即可。每次新的折痕都是 0 1 0 1……
func GetFolds(n int) []int {
	var res = make([]int, 0)
	return getFolds(1, n, res, 0)

}

func getFolds(level int, n int, res []int, x int) []int {
	if level > n {
		return res
	}

	res = getFolds(level+1, n, res, 0)
	res = append(res, x)
	res = getFolds(level+1, n, res, 1)
	return res
}

type SpecialTreeNode struct {
	Value  int
	Left   *SpecialTreeNode
	Right  *SpecialTreeNode
	Parent *SpecialTreeNode
}

func NewSpecialTreeNode(val int) *SpecialTreeNode {
	return &SpecialTreeNode{Value: val}
}

// GetSuccessor 获取任意一个节点node的后继节点
// 后继节点的定义：中序遍历的下一个节点
// 步骤：
// 1.如果节点有右子树，则后继节点是右子树的最左节点
// 2.如果节点没有右子树，则向上找父节点，直到某一节点是其父节点的左子节点，那个父节点就是后继节点（即这个节点是那个父节点左子树的最右节点）
func GetSuccessor(node *SpecialTreeNode) *SpecialTreeNode {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		cur := node.Right
		for cur != nil {
			if cur.Left == nil {
				return cur
			}
			cur = cur.Left
		}
	}

	cur := node
	parent := node.Parent
	for parent != nil {
		if cur == parent.Left {
			return parent
		}
		cur = parent
		parent = parent.Parent
	}
	return nil

}

// ConvertToDoubleNode 二叉树和双向链表在结构上有相似之处。
// 二叉树转换成双向链表的方式是按照中序遍历结果形成链表，left指针当作prev指针，right指针当作next指针
func ConvertToDoubleNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	start, _ := convertToDoubleNode(node)
	return start
}

// 输入一个二叉树的节点，返回它形成的双向链表的头节点和尾节点
// 一个节点左孩子形成的链表的end节点一定是这个节点左子树的最右节点；右孩子同理
func convertToDoubleNode(node *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}

	leftStart, leftEnd := convertToDoubleNode(node.Left)
	rightStart, rightEnd := convertToDoubleNode(node.Right)

	if leftEnd != nil {
		leftEnd.Right = node
	}
	if rightStart != nil {
		rightStart.Left = node
	}
	node.Left = leftEnd
	node.Right = rightStart

	start := leftStart
	end := rightEnd

	if start == nil {
		start = node
	}
	if end == nil {
		end = node
	}
	return start, end
}

// PrintEdge 逆时针获取边界节点
// 边界定义：
// a.头节点
// b.叶节点
// c.每层最左或最右节点
// 步骤：
// 1.记录每层最左和最右节点
// 2.从上到下打印所有层的最左节点
// 3.先序遍历，打印那些不属于某一层最左或最右的节点，但同时是叶节点的节点
// 4.从下到上打印所有层的最右节点，但节点不能同时是最左节点
func PrintEdge(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	height := getHeight(node, 0)
	var edgeTuple [][]*TreeNode
	for i := 0; i < height; i++ {
		eachLevel := make([]*TreeNode, 2)
		edgeTuple = append(edgeTuple, eachLevel)
	}
	setEdgeTuple(node, 0, edgeTuple)

	res := make([]int, 0)
	// 打印左边界
	for i := 0; i < height; i++ {
		res = append(res, edgeTuple[i][0].Value)
	}
	// 打印叶子节点
	res = printLeafNotInMap(node, 0, edgeTuple, res)
	// 打印右边界
	for i := height - 1; i >= 0; i-- {
		if edgeTuple[i][0] != edgeTuple[i][1] {
			res = append(res, edgeTuple[i][1].Value)
		}
	}
	return res
}

func printLeafNotInMap(node *TreeNode, level int, edgeTuple [][]*TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	if node.Left == nil && node.Right == nil && node != edgeTuple[level][0] && node != edgeTuple[level][1] {
		res = append(res, node.Value)
	}
	res = printLeafNotInMap(node.Left, level+1, edgeTuple, res)
	res = printLeafNotInMap(node.Right, level+1, edgeTuple, res)
	return res
}

func setEdgeTuple(node *TreeNode, level int, edgeTuple [][]*TreeNode) {
	if node == nil {
		return
	}

	// 按照递归顺序，先递归左节点
	// 因此每一层第一次给edgeTuple[level][0]设置值时设置的为最左节点
	// 每一层最后一次给edgeTuple[level][1]设置值时设置的为最右节点
	if edgeTuple[level][0] == nil {
		edgeTuple[level][0] = node
	}
	edgeTuple[level][1] = node
	setEdgeTuple(node.Left, level+1, edgeTuple)
	setEdgeTuple(node.Right, level+1, edgeTuple)
}

func getHeight(node *TreeNode, level int) int {
	if node == nil {
		return level
	}
	return int(math.Max(float64(getHeight(node.Left, level+1)), float64(getHeight(node.Right, level+1))))
}

// GetMaxLenWithSum 求二叉树中累加和为sum的最长路径
// 步骤：
// 1.用一个哈希表sumMap，记录从head开始的一条路径上的累加和出现的情况
//	 key代表某个累加和，value代表这个累加和在路径中最早出现的层数
//   如果在遍历到cur时，我们能够知道从head到cur路径上的累加和出现的情况，那么求以cur结尾的累加和为指定值的最长路径长度就非常容易
// 2.sumMap的更新过程：
//		1)首先加入(0,0)，表示累加和0不用包括任何节点就能得到
//		2)先序遍历，遍历到的节点为cur，从head到cur父节点的累加和为preSum，cur所在层数为level
//		3)将cur.value+preSum的值记为curSum，就是head到cur的累加和
//		4)如果sumMap中包含了curSum，就不更新，否则加入(curSum, level)
//		5)求解：在必须以cur结尾的情况下，累加和为规定值的最长路径长度
//		6)遍历cur的左子树和右子树，步骤同上
//		7)处理完以cur为头节点的子树，返回cur的父节点，在返回前在sumMap中查询curSum对应的level，如果等于当前level就删除记录
// 关键是要先序遍历，这样才能从上到下进行计算
func GetMaxLenWithSum(node *TreeNode, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 0
	level := 1
	preSum := 0
	maxLen := 0
	return preOrder(node, level, preSum, maxLen, k, sumMap)
}

func preOrder(node *TreeNode, level int, preSum int, maxLen int, k int, sumMap map[int]int) int {
	if node == nil {
		return maxLen
	}

	curSum := preSum + node.Value
	_, ok1 := sumMap[curSum]
	if !ok1 { // 如果sumMap中之前没有curSum这个key，说明level层是最早能使得从头节点开始累加得到curSum的层
		sumMap[curSum] = level
	}

	_, ok2 := sumMap[curSum-k]
	if ok2 {
		maxLen = int(math.Max(float64(level-sumMap[curSum-k]), float64(maxLen)))
	}

	maxLen = preOrder(node.Left, level+1, curSum, maxLen, k, sumMap)
	maxLen = preOrder(node.Right, level+1, curSum, maxLen, k, sumMap)

	// 遍历完以后，如果发现level层就是最早能使得从头节点开始累加得到curSum的层，那么删除这个key
	// 因为以这个节点开始向下遍历的过程已经结束，再去遍历其他节点时，不应该将这个节点造成影响带过去
	// 例如a节点的左孩子是b节点，右孩子是c节点，在遍历b节点及其子树时，在b节点记录了sumMap[curSum]=level
	// 那么，在离开b节点去遍历c节点时，应该将这个key删除掉
	// 如果sumMap中的值不是level，那就只可能比level小，说明是上层节点造成的影响，可以带到c节点的遍历中去
	if level == sumMap[curSum] {
		delete(sumMap, curSum)
	}
	return maxLen
}

// GetMaxBST 子树中最大的搜索二叉树
func GetMaxBST(node *TreeNode) *TreeNode {
	head, _, _, _ := getMaxBST(node)
	return head
}

func getMaxBST(node *TreeNode) (bstHead *TreeNode, bstSize int, min int, max int) {
	if node == nil {
		return nil, 0, math.MaxInt, math.MinInt
	}

	leftHead, leftSize, leftMin, leftMax := getMaxBST(node.Left)
	rightHead, rightSize, rightMin, rightMax := getMaxBST(node.Right)

	min = int(math.Min(float64(node.Value), math.Min(float64(leftMin), float64(rightMin))))
	max = int(math.Max(float64(node.Value), math.Max(float64(leftMax), float64(rightMax))))
	bstSize = int(math.Max(float64(leftSize), float64(rightSize)))
	if leftSize >= rightSize {
		bstHead = leftHead
	} else {
		bstHead = rightHead
	}
	if leftHead == node.Left && rightHead == node.Right && node.Value > leftMax && node.Value < rightMin {
		bstSize = leftSize + rightSize + 1
		bstHead = node
	}
	return bstHead, bstSize, min, max
}

// GetMaxBSTopoSize 二叉树中符合搜索二叉树条件的最大拓扑结构（不一定是完整的子树）
// 核心思路是后序遍历，对每个节点记录拓扑贡献(left, right)，表示这个节点左右子树中符合条件的最大size，并自底向上进行调整
func GetMaxBSTopoSize(node *TreeNode) int {
	node2Record := make(map[*TreeNode]*MaxBSTopoRecord)
	return postOrder(node, node2Record)
}

func postOrder(node *TreeNode, node2Record map[*TreeNode]*MaxBSTopoRecord) int {
	if node == nil {
		return 0
	}

	leftSize := postOrder(node.Left, node2Record)
	rightSize := postOrder(node.Right, node2Record)

	modifyMap(node.Left, node.Value, node2Record, true)
	modifyMap(node.Right, node.Value, node2Record, false)

	recordOfLeftTree := node2Record[node.Left]
	recordOfRightTree := node2Record[node.Right]

	var leftRecord int
	var rightRecord int
	if recordOfLeftTree != nil {
		leftRecord = recordOfLeftTree.Left + recordOfLeftTree.Right + 1
	}
	if recordOfRightTree != nil {
		rightRecord = recordOfRightTree.Left + recordOfRightTree.Right + 1
	}
	node2Record[node] = &MaxBSTopoRecord{
		Left:  leftRecord,
		Right: rightRecord,
	}
	return int(math.Max(float64(leftRecord+rightRecord+1), math.Max(float64(leftSize), float64(rightSize))))
}

func modifyMap(node *TreeNode, value int, node2Record map[*TreeNode]*MaxBSTopoRecord, b bool) int {
	_, ok := node2Record[node]
	if node == nil || !ok {
		return 0
	}

	record := node2Record[node]
	if (b && node.Value > value) || (!b && node.Value < value) {
		delete(node2Record, node)
		return record.Left + record.Right + 1
	} else {
		var child *TreeNode
		if b {
			child = node.Right
		} else {
			child = node.Left
		}
		minus := modifyMap(child, value, node2Record, b)
		if b {
			record.Right = record.Right - minus
		} else {
			record.Left = record.Left - minus
		}
		node2Record[node] = record
		return minus
	}

}

type MaxBSTopoRecord struct {
	Left  int
	Right int
}

func TraverseZigZag(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	res := make([]int, 0)
	deque := basic.NewDeque()
	last := node
	var nextLast *TreeNode
	var cur *TreeNode
	deque.AddFirst(node)
	flag := true
	for !deque.IsEmpty() {
		if flag {
			cur = deque.PollFirst().(*TreeNode)
			res = append(res, cur.Value)
			if cur.Left != nil {
				deque.AddLast(cur.Left)
				if nextLast == nil {
					nextLast = cur.Left
				}
			}
			if cur.Right != nil {
				deque.AddLast(cur.Right)
				if nextLast == nil {
					nextLast = cur.Right
				}

			}
		} else {
			cur = deque.PollLast().(*TreeNode)
			res = append(res, cur.Value)
			if cur.Right != nil {
				deque.AddFirst(cur.Right)
				if nextLast == nil {
					nextLast = cur.Right
				}
			}
			if cur.Left != nil {
				deque.AddFirst(cur.Left)
				if nextLast == nil {
					nextLast = cur.Left
				}
			}
		}

		if cur == last && !deque.IsEmpty() {
			last = nextLast
			nextLast = nil
			flag = !flag
		}
	}

	return res
}

// FindTwoErrorNodes 搜索二叉树中有两个节点位置错误，调换它们位置后重新形成搜索二叉树，找出这两个节点
// 中序遍历二叉树，会出现两种情况：
// 1.有两次降序，第一次降序的第一个节点和第二次降序的第二个节点就是那两个错误节点
// 2.有一次降序，这次降序的第一个节点和第二个节点就是那两个错误节点
// 综合来看，就是第一次降序的第一个节点和最后一次降序的第二个节点
func FindTwoErrorNodes(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}

	res := make([]*TreeNode, 2)
	stack := basic.NewStack()
	stack.Push(node)
	cur := node.Left

	var prev *TreeNode
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}

		cur = stack.Pop().(*TreeNode)
		if prev != nil {
			if cur.Value < prev.Value {
				if res[0] == nil {
					res[0] = prev
				}
				res[1] = cur
			}
		}
		prev = cur
		cur = cur.Right
	}

	return res
}

// Contains node1是否包含整个node2拓扑结构（node2未必是node1中一棵完整的子树）
func Contains(node1 *TreeNode, node2 *TreeNode) bool {
	if node2 == nil {
		return true
	}
	if node1 == nil {
		return false
	}
	return check(node1, node2) || Contains(node1.Left, node2) || Contains(node1.Right, node2)
}

func check(node1 *TreeNode, node2 *TreeNode) bool {
	if node2 == nil {
		return true
	}
	if node1 == nil || node1.Value != node2.Value {
		return false
	}
	return check(node1.Left, node2.Left) && check(node1.Right, node2.Right)
}

func BuildBSTByPostOrder(arr []int) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	return buildBSTByPostOrder(arr, 0, len(arr)-1)
}

func buildBSTByPostOrder(arr []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	head := NewTreeNode(arr[end])
	less := math.MinInt
	more := math.MaxInt
	for i := start; i < end; i++ {
		if arr[end] > arr[i] {
			less = i
		} else {
			if more == math.MaxInt {
				more = i
			}
		}
	}
	head.Left = buildBSTByPostOrder(arr, start, less)
	head.Right = buildBSTByPostOrder(arr, more, end-1)
	return head
}

// BuildBalancedBSTBySortedArr 根据有序数组生成平衡搜索二叉树
func BuildBalancedBSTBySortedArr(arr []int) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	return buildBalancedBSTBySortedArr(arr, 0, len(arr)-1)
}

func buildBalancedBSTBySortedArr(arr []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	head := NewTreeNode(arr[mid])
	head.Left = buildBalancedBSTBySortedArr(arr, start, mid-1)
	head.Right = buildBalancedBSTBySortedArr(arr, mid+1, end)
	return head
}

func GetMaxDistance(node *TreeNode) int {
	maxDistance, _ := getMaxDistance(node)
	return maxDistance
}

func getMaxDistance(node *TreeNode) (maxDistance int, height int) {
	if node == nil {
		return 0, 0
	}
	leftMaxDistance, leftHeight := getMaxDistance(node.Left)
	rightMaxDistance, rightHeight := getMaxDistance(node.Right)
	height = int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	maxDistance = int(math.Max(math.Max(float64(leftMaxDistance), float64(rightMaxDistance)), float64(leftHeight+rightHeight+1)))
	return maxDistance, height
}
