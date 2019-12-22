/*
 * // This is the custom function interface.
 * // You should not implement it, or speculate about its implementation
 * class CustomFunction {
 * public:
 *     // Returns f(x, y) for any given positive integers x and y.
 *     // Note that f(x, y) is increasing with respect to both x and y.
 *     // i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 *     int f(int x, int y);
 * };
 */
// 二分：O(xlogy)
class Solution {
public:
    vector<vector<int>> findSolution(CustomFunction& customfunction, int z) {
        vector<vector<int>> ans;
    
        for(int y = 1; y <= 1000; y++) {
            int xlow = 1, xhigh = 1000;
            while(xlow <= xhigh) { //
                int xmid = xlow + ((xhigh - xlow) >> 1); // 优先级：`+`大于`>>`
                int cal = customfunction.f(xmid, y);
                if(cal < z) {
                    xlow = xmid + 1;
                } else if(cal == z) {
                    ans.push_back({xmid, y});
                    break;
                } else {
                    xhigh = xmid - 1;
                }
                // cout << "cal: " << cal << " xmid: " << xmid <<" xlow: " << xlow << " xhigh: " << xhigh << '\n';
            }
        }

        return ans;
    }
};

// 同时利用f(x, y)的两个特性，O(x + y)，函数双指针
class Solution {
public:
    vector<vector<int>> findSolution(CustomFunction& customfunction, int z) {
        vector<vector<int>> ans;
        
        int x = 1, y = 1000;
        while(x <= 1000 && y >= 1) {
            int cal = customfunction.f(x, y);
            if(cal > z) {
                y--;
            } else if(cal < z) {
                x++;
            } else {
                ans.push_back({x++, y--});
            }
        }

        return ans;
    }
};