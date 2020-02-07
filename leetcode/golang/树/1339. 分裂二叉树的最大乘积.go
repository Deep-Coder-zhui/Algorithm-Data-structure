/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
子树之和尽可能接近二叉树的结点之和的一半
时间：O(n)，空间：O(d)，n为树结点数，d为树深
*/
var (
    p = 1000000007
    part = 0
    sum = 0
)

func getSum(root *TreeNode) {
    if root != nil {
        sum += root.Val
        getSum(root.Left)
        getSum(root.Right)
    }
}

func getPartSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := getPartSum(root.Left), getPartSum(root.Right)
    if l != 0 && int(math.Abs(float64(l-sum/2))) < int(math.Abs(float64(part-sum/2))) {
        part = l
    }
    if r != 0 && int(math.Abs(float64(r-sum/2))) < int(math.Abs(float64(part-sum/2))) {
        part = r
    }
    return l+r+root.Val
}

func maxProduct(root *TreeNode) int {
    sum, part = 0, 0 // 每次必须要初始化
    getSum(root)
    getPartSum(root)
    // fmt.Println(part, sum-part)
    return int(int64(part)*int64(sum-part)%int64(p))
}

/*
后序遍历计算以每个结点为根的子树的结点之和，较前面方法的优化在于把一次递归计算子树结点之和变为循环遍历
时间：O(n)，空间：O(n)，n为树结点数
*/
const p = 1000000007
var sum []int

func getSum(root *TreeNode) int { // 本质是后序遍历
    if root == nil {
        return 0
    }
    sum = append(sum, root.Val + getSum(root.Left) + getSum(root.Right))
    return sum[len(sum)-1]
}

func maxProduct(root *TreeNode) int {
    sum = []int{} //
    getSum(root)
    
    ans := int64(0)
    for i := 0; i < len(sum); i++ {
        ans = max(ans, int64(sum[i])*int64(sum[len(sum)-1]-sum[i]))
    }
    
    return int(ans%p)
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}
