/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    
    // KMP算法实现
    return kmpSearch(haystack, needle)
}

// 构建部分匹配表
func buildPartialMatchTable(pattern string) []int {
    m := len(pattern)
    lps := make([]int, m) // longest proper prefix which is also suffix
    length := 0
    i := 1
    
    for i < m {
        if pattern[i] == pattern[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                length = lps[length-1]
            } else {
                lps[i] = 0
                i++
            }
        }
    }
    
    return lps
}

// KMP搜索算法
func kmpSearch(text string, pattern string) int {
    n := len(text)
    m := len(pattern)
    
    if m > n {
        return -1
    }
    
    lps := buildPartialMatchTable(pattern)
    i := 0 // text的索引
    j := 0 // pattern的索引
    
    for i < n {
        if text[i] == pattern[j] {
            i++
            j++
            
            if j == m {
                return i - j // 找到匹配
            }
        } else {
            if j != 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    
    return -1 // 未找到匹配
}

// 暴力匹配算法（作为对比）
func bruteForceSearch(text string, pattern string) int {
    n := len(text)
    m := len(pattern)
    
    for i := 0; i <= n-m; i++ {
        found := true
        for j := 0; j < m; j++ {
            if text[i+j] != pattern[j] {
                found = false
                break
            }
        }
        if found {
            return i
        }
    }
    
    return -1
}

// @lc code=end

// 测试用例
/*
func main() {
    // 测试KMP算法
    fmt.Println(strStr("hello", "ll")) // 2
    fmt.Println(strStr("aaaaa", "bba")) // -1
    fmt.Println(strStr("", "")) // 0
    fmt.Println(strStr("abc", "c")) // 2
    fmt.Println(strStr("mississippi", "issip")) // 4
}
*/