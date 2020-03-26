#include<iostream>
#include<string>
#include<vector>
using namespace std;

class Solution {
public:
    /*
    * @param s: the IP string
    * @return: a valid IP address
    */
    string validIpAddress(string &s) {
        if (s.size() < 4 || s.size() > 12)
            return "";
        int len = s.size();
        for (int i = 1; i <= 3; ++i)
        {
            for (int j = 1; j <= 3; ++j)
            {
                for (int k = 1; k <= 3; ++k)
                {
                    string str = s;
                    if (i + j + k < len && isValid(i,j,k,str))
                    {
                        str.insert(i, ".");
                        str.insert(i + j + 1, ".");
                        str.insert(i + j + k + 2, ".");
                        return str; // 如果要求所有情形，可用vector保存结果
                    }
                }
            }
        }
        return ""; // 没有合法ip
    }
    
    bool isValid(int i, int j, int k, string str)
    {
        int len = str.size();
        string si = str.substr(0, i); // 截取下标为0，长度为i的字符串
        string sj = str.substr(i,j);
        string sk = str.substr(i + j, k);
        string sd = str.substr(i + j + k, len - i - j - k);
        if (stoi(si) <= 255 && stoi(sj) <= 255 && stoi(sk) <= 255 && stoi(sd) <= 255)
        {
            // 任何一个以点号隔开的ip整数表示值大于0时不能以0开头，非法:10.08.1.1
            if ((si[0] == '0'&&i > 1) || (sj[0] == '0'&&j > 1) || (sk[0] == '0'&&k > 1) || (sd[0] == '0'&&len - i - j - k > 1))
                return false;
            else
                return true;
        }
        else
            return false;
    }
};

int main() {
    Solution sol;
    string s;
    while(cin >> s) {
        cout << sol.validIpAddress(s) << endl;
    }
}
