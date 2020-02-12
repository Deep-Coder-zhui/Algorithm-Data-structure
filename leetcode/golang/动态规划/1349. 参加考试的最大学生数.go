/*
方法1：状态压缩DP
refer: https://leetcode-cn.com/problems/maximum-students-taking-exam/solution/zhuang-tai-ya-suo-dp-by-lucifer1004/
oneCnt[state]: 当某行状态为state时，二进制表示中1的个数，也即为学生入座的人数
initState[r]：第r行的初始状态，二进制表示中1代表一定不能入座，0表示可能入座
dp[r][state] = max{dp[r][state], dp[r-1][_state] + oneCnt[state]}，dp[r][state]: 第r行状态为state时前r行最多入座学生，下标从1开始；
state需要满足：
1、本行：不与初始座位分布冲突，入座学生左右侧不可坐人 
2、与上一行：上一行座位选择不与初始座位分布冲突，本行坐人的位置左上和右上不可坐人
最终答案为max{dp[maxrow][state]}
时间：O(m*4^n)，空间：O(n*2^n)
方法2：二分图最大独立集
https://leetcode-cn.com/problems/maximum-students-taking-exam/solution/er-fen-tu-zui-da-du-li-ji-by-lightcml/
??
*/
func maxStudents(seats [][]byte) int {
    oneCnt, initState := [256]int{}, [9]int{}
    dp := [9][256]int{}
    
    for state := 0; state < 1<<len(seats[0]); state++ {
        oneCnt[state] = oneCnt[state>>1] + (state&1) // trick
    }
    for r := 0; r < len(seats); r++ {
        for c := 0; c < len(seats[r]); c++ {
            if seats[r][c] == '#' {
                initState[r+1] |= (1<<c) // 第c列为1
            }
        }
    }
    
    for r := 0; r < len(seats); r++ {
        for state := 0; state < 1<<len(seats[0]); state++ {
            if state&initState[r+1]==0 && state&(state>>1)==0 && state&(state<<1)==0 { // go语言中&和<<,>>操作符同级，高于==
                for _state := 0; _state < 1<<len(seats[0]); _state++ {
                    if _state&initState[r]==0 && state&(_state>>1)==0 && state&(_state<<1)==0 {
                        dp[r+1][state] = max(dp[r+1][state], dp[r][_state] + oneCnt[state])
                    }
                }
            }
        }
    }
    
    maxStu := 0
    for state := 0; state < 1<<len(seats[0]); state++ {
        maxStu = max(maxStu, dp[len(seats)][state])
    }
    
    return maxStu
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
