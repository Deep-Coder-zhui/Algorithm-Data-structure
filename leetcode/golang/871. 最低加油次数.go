/*
sour--->stations[0]......stations[n-1]--->dest
0-1背包，加油站从0算起
interval[i]: 第i个加油站到前面一个站（i=0，则为起点）的距离
dp[i][j]：到第i个加油站且共加j(最大为i+1)次油后最大还能行驶的距离（dp[i][j]=-1表示不可达）
dp[i][j] = max{dp[i-1][j] - interval[i], dp[i-1][j-1] - interval[i][0] + stations[i][1]}
*/
const stations_len = 505
func minRefuelStops(target int, startFuel int, stations [][]int) int {
    interval := [stations_len]int{}
    if len(stations) == 0 { //
        if startFuel < target {
            return -1
        }
        return 0
    }
    interval[0] = stations[0][0]
    for i := 1; i < len(stations); i++ {
        interval[i] = stations[i][0] - stations[i-1][0]
    }
    
    dp := [stations_len][stations_len]int{}
    for i := 0; i < len(stations); i++ {
        for j := 0; j < len(stations); j++ {
            dp[i][j] = -1 //
        }
    }
    if startFuel < interval[0] {
        return -1
    } 
    dp[0][0] = startFuel - interval[0]
    if dp[0][0] >= target - stations[0][0] {
        return 0
    }
    dp[0][1] = dp[0][0] + stations[0][1]
    if dp[0][1] >= target - stations[0][0] {
        return 1
    }

    for i := 1; i < len(stations); i++ {
        for j := 0; j <= i+1; j++ {
            if j == 0 {
                if dp[i-1][j] - interval[i] >= 0 {
                    dp[i][j] = dp[i-1][j] - interval[i]
                }                
            } else {
                if dp[i-1][j-1] - interval[i] >= 0 { //
                    dp[i][j] = dp[i-1][j-1] - interval[i] + stations[i][1]
                }
                if dp[i-1][j] - interval[i] >= 0 && dp[i-1][j] - interval[i] > dp[i][j] {
                     dp[i][j] = dp[i-1][j] - interval[i]
                }
            }
        }
    }
    
    for j := 0; j <= len(stations); j++ { //
        if dp[len(stations)-1][j] != -1 && dp[len(stations)-1][j] >= target - stations[len(stations)-1][0] {
            return j
        }
    }
    return -1
}