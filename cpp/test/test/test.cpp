
#include <iostream>
#include <unordered_map>
#include <vector>
#include <list>
#include <set>
#include <functional>
#include <algorithm>
#include <stack>

using namespace std;

class LRU {
    int m_capacity;
    typedef pair<int, int> node;
    list <node> m_data;
    unordered_map<int, list<node>::iterator> m_hash;

public:
    LRU(int capacity) {
        m_capacity = capacity;
        m_data.clear();
        m_hash.clear();
    }

    int get(int key) {
        auto it = m_hash.find(key);
        if (it == m_hash.end()) {
            return -1;
        }
        node t(key, it->second->second);
        m_data.erase(it->second);
        m_data.push_front(t);
        m_hash[key] = m_data.begin();
        return t.second;
    }

    void put(int key, int v) {
        node t(key, v);
        auto it = m_hash.find(key);
        if (it != m_hash.end()) {
            m_data.erase(it->second);
            m_data.push_front(t);
            m_hash[key] = m_data.begin();
            return;
        }
        if (m_capacity == m_data.size()) {
            m_hash.erase(m_data.end()->first);
            m_data.pop_back();
        }
        m_data.push_front(t);
        m_hash[key] = m_data.begin();
    }
};

double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
    int m = nums1.size();
    int n = nums2.size();
    int i = 0, j = 0, m1 = 0, m2 = 0;

    // 查找中位数
    for (int count = 0; count <= (m + n) / 2; ++count) {
        // 如果两数组长度为偶数位，则保留2位
        m2 = m1;
        if (i < n && j < m) {
            if (nums1[i] > nums2[j]) {
                m1 = nums2[j++];
            }
            else {
                m1 = nums1[i++];
            }
        }
        else if (i < m) {
            m1 = nums1[i++];
        }
        else {
            m1 = nums2[j++];
        }
    }

    if ((m + n) % 2 == 1) {
        return static_cast<double>(m1);
    }
    return (static_cast<double>(m1) + static_cast<double>(m2)) / 2.0;
}

string convert(string s, int numRows) {
    int n = s.length();
    if (numRows <= 1 || n <= numRows) return s;

    int cycle = (numRows - 1) << 1;  // 下一个字符的间隔数
    string ans(n, '\0');
    int idx = 0;

    for (int r = 0; r < numRows; r++) {
        for (int i = r; i < n; i += cycle) {
            ans[idx++] = s[i];  // 按周期更新每行的字符
            if (r > 0 && r < numRows - 1) {
                int j = i + cycle - (r << 1);
                if (j < n) {
                    ans[idx++] = s[j];
                }
            }
        }
    }

    return ans;
}

bool isMatch(string s, string p) {
    int n = s.length(), m = p.length();
    vector<vector<bool>> memo(n + 1, vector<bool>(m + 1, false));
    function<bool(int, int)> dp = [&](int i, int j)->bool {
        if (memo[i][j]) {
            return memo[i][j];
        }

        if (j == m) {
            return i == n;
        }

        bool first = false;
        if (i < n && (p[j] == s[i] || p[j] == '.')) {
            first = true;
        }
        bool ans = false;
        if (j <= (m - 2) && p[j + 1] == '*') {
            ans = (dp(i, j + 2) || (first && dp(i + 1, j)));
        }
        else {
            ans = (first && dp(i + 1, j + 1));
        }
        memo[i][j] = ans;
        return ans;
        };
    return dp(0, 0);
}

int romanToInt(string s) {
    string symbol[] = { "M", "D", "C", "L", "X", "V", "I" };
    string symbolCal[] = { "CM", "CD", "XC", "XL", "IX", "IV" };
    int val[] = { 1000, 500, 100, 50, 10, 5, 1 };
    int valCal[] = { 900, 400, 90, 40, 9, 4 };
    int ans = 0;
    int l = s.length();
    int pos = 0;
    int reset = false;
    while (pos < l) {
        for (int i = 0; i < sizeof(symbolCal) / sizeof(symbolCal[0]); i++) {
            cout << i <<  " " << symbolCal[i] << " " << s.substr(pos, 2) << endl;
            if (pos + 1 < l && s.substr(pos, 2) == symbolCal[i]) {
                ans += valCal[i];
                pos += 2;
                reset = true;
            }
        }

        if (reset) {
            reset = false;
            continue;
        }

        for (int i = 0; i < sizeof(symbol) / sizeof(symbol[0]); i++) {
            cout << i << " " << symbol[i] << " " << s.substr(pos, 1) << endl;
            if (pos < l && s.substr(pos, 1) == symbol[i]) {
                ans += val[i];
                pos++;
                break;
            }
        }
    }
    return ans;
}

string longestCommonPrefix(vector<string>& strs) {
    string ans = "";
    int index = 0;
    int size = strs.size();
    bool end = false;
    while (1) {
        char c = 0;
        for (int i = 0; i < size; i++) {
            const string& s = strs[i];
            if (s.length() <= index) {
                end = true;
                break;
            }
            if (c == 0) {
                c = s[index];
            }

            if (s[index] != c) {
                end = true;
                break;
            }
        }

        if (end) {
            break;
        }
        ans += c;
        index++;
    }
    return ans;
}

vector<vector<int>> threeSum(vector<int>& nums) {
    vector<vector<int>> ans;
    sort(nums.begin(), nums.end());

    for (int i = 0; i < nums.size(); i++) {
        if (i > 0 && nums[i] == nums[i - 1]) {
            continue;
        }
        int j = i + 1;
        int k = nums.size() - 1;
        while (j < k) {
            int total = nums[i] + nums[j] + nums[k];

            if (total > 0) {
                k--;
            }
            else if (total < 0) {
                j++;
            }
            else {
                ans.push_back({ nums[i], nums[j], nums[k] });
                j++;
                while (nums[j] == nums[j - 1] && j < k) {
                    j++;
                }
            }
        }
    }

    return ans;
}

vector<int> findSubstring(string s, vector<string>& words) {
    vector<int> ans;
    if (s.length() == 0 || words.empty()) {
        return ans;
    }
    unordered_map<string, int> hash;
    int l = words[0].size();
    int total = l * words.size();
    for (int i = 0; i < words.size(); i++) {
        hash[words[i]]++;
    }

    int index = 0, n = s.length();
    while (index + total <= n) {
        bool find = true;
        unordered_map<string, int> t;
        int start = index;
        while (start + l <= n && start < index + total) {
            string subS = s.substr(start, l);
            auto it = hash.find(subS);
            if (it == hash.end()) {
                find = false;
                break;
            }
            t[subS]++;
            if (t[subS] > hash[subS]) {
                find = false;
                break;
            }
            start += l;
        }

        if (find) {
            ans.push_back(index);
        }
        index++;
    }
    return ans;
}

int validParentheses(string& s, int& i) {
    stack<char> st;
    int count = 0;
    st.push(s[i]);
    int n = s.length();
    i = i + 1;
    while (i < n) {
        if (!st.empty() && s[i] == ')') {
            count += 2;
            st.pop();
        } else {
            if (s[i] == '(') {
                st.push(s[i]);
            }
            if (st.empty()) {
                break;
            }
        }
        i++;
    }
    return count;
}

int longestValidParentheses(string s) {
    if (s.length() <= 0) {
        return 0;
    }
    int i = 0;
    int n = s.length();
    int ans = 0;
    while (i < n) {
        if (s[i] != '(') {
            i++;
            continue;
        }

        int sub = validParentheses(s, i);

        ans = max(ans, sub);
    }
    return ans;
}

void printMa(vector<vector<int>>& vv) {
    for (auto& it : vv) {
        for (auto& i : it) {
            cout << i << " ";
        }
        cout << endl;
    }

    cout << endl;
}

vector<vector<int>> generateMatrix(int n) {
    vector<vector<int>> ans(n, vector<int>(n, 0));

    int count = 1;
    int l = 0, r = n - 1, top = 0, bottom = n - 1;
    while (l <= r && top <= bottom) {
        for (int i = l; i <= r; i++) {
            ans[top][i] = count++;
        }
        top++;

        printMa(ans);

        for (int i = top; i <= bottom; i++) {
            ans[i][r] = count++;
        }
        r--;
        printMa(ans);

        if (l <= r) {
            for (int i = r; i >= l; i--) {
                ans[bottom][i] = count++;
            }
            bottom--;
        }
        printMa(ans);

        if (top <= bottom) {
            for (int i = bottom; i >= top; i--) {
                ans[i][l] = count++;
            }
            l++;
        }
        printMa(ans);
    }
    return ans;
}

void p(vector<int>& v) {
    for (int i : v) {
        cout << i << " ";
    }
    cout << endl;
}

string getPermutation(int n, int k) {
    vector<int> fact(n + 1, 1);
    for (int i = 1; i <= n; i++) {
        fact[i] = fact[i - 1] * i;
    }

    p(fact);
    string ans;
    vector<bool> used(n + 1);
    for (int i = 0; i < n; i++) {
        int cnt = fact[n - i - 1];
        cout << "cnt " << cnt << " ";
        for (int j = 1; j <= n; j++) {
            if (used[j]) continue;

            if (k > cnt) k -= cnt;
            else {
                used[j] = true;
                ans += j + '0';
                break;
            }
        }
    }
    return ans;
}

int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
    int m = obstacleGrid.size();
    int n = obstacleGrid[0].size();
    vector<vector<int>> dp(m, vector<int>(n, 1));

    for (int i = 1; i < m; i++) {
        for (int j = 1; j < n; j++) {
            if (obstacleGrid[i][j] != 1 && obstacleGrid[i][j - 1] != 1 && obstacleGrid[i - 1][j] != 1) {
                dp[i][j] = dp[i][j - 1] + dp[i - 1][j];
                continue;
            }
            
            if (obstacleGrid[i][j] == 1) {
                dp[i][j] = 0;
            }

            if (obstacleGrid[i][j - 1] == 1) {
                dp[i][j - 1] = 0;
            }

            if (obstacleGrid[i - 1][j] == 1) {
                dp[i - 1][j] = 0;
            }
        }
    }
    return dp[m - 1][n - 1];
}

bool isNumber(string s) {
    int n = s.length();
    bool isdot = false, nums = false, ise = false;
    for (int i = 0; i < n; i++) {
        if (isdigit(s[i])) nums = true;
        else if (s[i] == '+' || s[i] == '-') {
            if (i > 0 && (s[i - 1] != 'e' && s[i - 1] != 'E')) return false;
        }
        else if (s[i] == 'e' || s[i] == 'E') {
            if (ise || !nums) return false;
            ise = true;
            nums = false;
        }
        else if (s[i] == '.') {
            if (isdot || ise) return false;
            isdot = true;
        }
        else {
            return false;
        }
    }
    return nums;
}

void testVectorIt() {
    vector<int> v{ 1,2,3,4,5,6 };

    for (auto it = v.begin(); it != v.end();) {
        if (*it > 2) {
            it = v.erase(it);
        }
        else {
            it++;
        }
    }

    for (auto it = v.begin(); it != v.end(); it++) {
        cout << *it << " ";
    }
    cout << endl;
}

int main()
{
    testVectorIt();
    //LRU l(3);
    //l.put(1, 1);
    //l.put(1, 1);
    //l.put(1, 1);
    //cout << l.get(1) << endl;
    //vector<int> v1{1,3};
    //vector<int> v2{ 2 };
    //cout << findMedianSortedArrays(v1, v2) << endl;
    /*string s{ "PAYPALISHIRING" };
    convert(s, 3);*/
    //string s{ "MCMXCIV" };
    //romanToInt(s);
    //vector<string> strs{ "flower","flow","flight" };
    //longestCommonPrefix(strs);
    //vector<int> v{ -2, -1, -1, 0, 2 };
    //threeSum(v);
    //string s = "wordgoodgoodgoodbestword";
    //vector<string> words = { "word","good","best","good" };
    //findSubstring(s, words);
    //string s = "(()";
    //cout << longestValidParentheses(s) << endl;
    //auto ans = generateMatrix(3);
    //printMa(ans);
    /*cout << getPermutation(9, 362770) << endl;*/
    //vector<vector<int>> v{ {0,1}, {0, 0} };
    //uniquePathsWithObstacles(v);
    /*isNumber("--6");*/

}

