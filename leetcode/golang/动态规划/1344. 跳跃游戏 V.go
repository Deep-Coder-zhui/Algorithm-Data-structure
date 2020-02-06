/*
注意到下标之间跳跃时条件：1、步伐d之内 2、对应数值更小。容易想到先对arr中结点node{数值, 下标}按数值升序，对node有：
dp[i] = max{dp[k]} + 1，0<=k<i & abs(结点i下标-结点k下标)<=d, dp[i]：从node[i]跳跃到前面的node的最大访问下标个数，当然，首先需要确保能够从i跳转到k
时间：O(n^3)(?)，空间：O(n)，n为数组大小
*/
type node struct {
    v, pos int
}

func maxJumps(arr []int, d int) int {
    tmp := []node{}
    for i, v := range arr {
        tmp = append(tmp, node{v, i})
    }
    sort.Slice(tmp, func(i, j int) bool {
        if tmp[i].v < tmp[j].v {
            return true
        }
        return false
    })
    
    dp := make([]int, len(tmp)) 
    for i := 0; i < len(dp); i++ {
        dp[i] = 1 //
    }
    
    ans := dp[0]
    for i := 1; i < len(tmp); i++ {
        for k := 0; k < i; k++ {
            if int(math.Abs(float64(tmp[i].pos - tmp[k].pos))) > d || tmp[i].v == tmp[k].v { // 值不能相等
                continue
            }
            canJump := true
            up := max(tmp[i].pos, tmp[k].pos)
            for t := min(tmp[i].pos, tmp[k].pos)+1; t < up; t++ {
                if arr[t] >= tmp[i].v {
                    canJump = false
                    break
                }
            }
            if canJump {
                dp[i] = max(dp[i], dp[k]+1)
            } 
        }
        ans = max(ans, dp[i])
    }
    
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
/*
优化：dp[i] = max{dp[i可移动到的坐标]} + 1，dp[i]表示从下标i出发可以获得的最大访问下标个数
时间：O(n^2)，空间：O(n)，n为数组大小
*/
type node struct {
    v, pos int
}

func maxJumps(arr []int, d int) int {
    tmp := []node{}
    for i, v := range arr {
        tmp = append(tmp, node{v, i})
    }
    sort.Slice(tmp, func(i, j int) bool {
        if tmp[i].v < tmp[j].v {
            return true
        }
        return false
    })
    
    dp := make([]int, len(tmp)) 
    ans := 0
    for i := 0; i < len(tmp); i++ {
        dp[tmp[i].pos] = 1
        // 向左跳跃
        bottom := max(0, tmp[i].pos-d)
        for k := tmp[i].pos-1; k >= bottom; k-- {
            if arr[k] >= tmp[i].v {
                break
            }
            if dp[k] != 0 {
                dp[tmp[i].pos] = max(dp[tmp[i].pos], dp[k]+1)
            }
        }
        // 向右跳跃
        up := min(len(tmp)-1, tmp[i].pos+d)
        for k := tmp[i].pos+1; k <= up; k++ {
            if arr[k] >= tmp[i].v {
                break
            }
            if dp[k] != 0 {
                dp[tmp[i].pos] = max(dp[tmp[i].pos], dp[k]+1)
            }
        }
        ans = max(ans, dp[tmp[i].pos])
    }
    
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
