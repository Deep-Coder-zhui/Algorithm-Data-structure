/*
Date: 2020-3-26
荷兰国旗问题
现有若干由红、白、蓝三种颜色的条块序列，要将它们重新排列使所有相同颜色的条块在一起，这里希望将所有红色的条块放在最左边，
所有白色条块放在中间，所有蓝色条块放在最右边。

refer: https://www.cnblogs.com/liuzhen1995/p/6439429.html
红白蓝分别用0、1、2表示
思路：双指针。定义三个指针，初始化begin=0, cur=0（可能指向元素为2）, end=len(goal)-1，然后cur右移：
if cur->0，则cur和begin指向元素完成交换，同时begin和cur右移
if cur->1，则cur右移
if cur->2，则cur和end指向元素完成交换，同时cur指针不动（可能交换前end所指元素为0），end左移
循环下去，直到cur > end，操作结束！
*/
package main

import "fmt"

func threeColors(goal []int) {
    if len(goal) < 2 {
        return
    }
    begin, cur, end := 0, 0, len(goal)-1
    for cur <= end {
        if goal[cur] == 0 {
            goal[begin], goal[cur] = goal[cur], goal[begin]
            begin++
            cur++
        } else if goal[cur] == 1 {
            cur++
        } else {
            goal[end], goal[cur] = goal[cur], goal[end]
            end--
        }
    }
}

func main() {
    // goal := []int{1, 2, 0, 1, 0, 2, 1, 0, 2}
    goal := []int{2, 1, 0, 1, 2}
    threeColors(goal)
    fmt.Println(goal)
}
