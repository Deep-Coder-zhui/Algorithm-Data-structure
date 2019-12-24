/*
显然，maxSize对本题来说是多余的
直接用unordered_map对符合要求的进行统计
40ms（beat 98.12% cpp submission）+ 15.8 MB
时间：O(minSize*(len(s)-minSize+1))，空间：O(minSize*(len(s)-minSize+1))
*/
class Solution {
public:
    int maxFreq(string s, int maxLetters, int minSize, int maxSize) {
        if(minSize > s.length()) {
            return 0;
        }
        
        unordered_map<string, int> umap;
        int cnt[26] = {0};
        int letters = 0;
        int start = 0, end = minSize-1;
        
        for(int i = start; i <= end; i++) {
            if(cnt[s[i]-'a']++ == 0) {
                ++letters;
            }
        }
        if(letters <= maxLetters) {
            umap[s.substr(start, end+1-start)]++;
        }
        
        for(int i = end+1; i < s.length(); i++) {
            if(--cnt[s[start++]-'a'] == 0) {
                letters--;
            }
            if(cnt[s[i]-'a']++ == 0) {
                letters++;
            }
            if(letters <= maxLetters) { //
                umap[s.substr(start, i+1-start)]++;
            }
        }
        
        int ans = 0;
        for(auto it = umap.begin(); it != umap.end(); it++) {
            if(it->second > ans) {
                ans = it->second;
            }
        }
        return ans;
    }
};