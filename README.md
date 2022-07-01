# golang_word_break

Given a string `s` and a dictionary of strings `wordDict`, return `true` if `s` can be segmented into a space-separated sequence of one or more dictionary words.

**Note** that the same word in the dictionary may be reused multiple times in the segmentation.

## Examples

**Example 1:**

```
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

```

**Example 2:**

```
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

```

**Example 3:**

```
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false

```

**Constraints:**

- `1 <= s.length <= 300`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 20`
- `s` and `wordDict[i]` consist of only lowercase English letters.
- All the strings of `wordDict` are **unique**.

## 解析

題目給定一個字串 s, 還有一個字串陣列 wordDict

要求寫出一個演算法來判斷 s 能不能夠由 wordDict 裏面的字串所b

其中可以重複使用 wordDict 裡面的字串

舉例來思考

假設給定 s: “leetcode”, wordDict : [”leet”, “code”]

從 start：0 

開始找尋可以組成的 dictWord

會有以下決策樹

![](https://i.imgur.com/eQFW6zk.png)

可以發現 

把從 i 開始到 s 結尾可以組成的值 表示為 $wordComposeFrom(i)$

可以注意到 以下關係式

$wordComposeFrom(i) = wordComposeFrom(i+len(word_k))$  if  $s[i:i+len(word_k)]==word_k$  

為了避免重複運算

所以可以倒過來運算

初始化 $wordComposeFrom(sLen) = true$b0

初始化 wordComposeFrom(i) = false, i = 0~ sLen-1

具體如下圖結構

![](https://i.imgur.com/kqA6DBj.png)

對於每個起始點 start = sLen-1 到 0 共有 sLen 個

要比對的 dictWord 有 len(dictWord) 個

所以時間複雜度是 O(n*m) , n 是 s 字串長度 , m 是字典字串個數

空間複雜度是 O(n) , n 是 s字串長度 

## 程式碼
```go
package sol

package sol

func wordBreak(s string, wordDict []string) bool {
	sLen := len(s)
	dp := make([]bool, sLen+1)
	// init dp[sLen] = true
	dp[sLen] = true
	for start := sLen - 1; start >= 0; start-- {
		for _, word := range wordDict {
			if start+len(word) <= sLen && s[start:start+len(word)] == word {
				dp[start] = dp[start+len(word)]
				break
			}
		}
	}
	return dp[0]
}
```
## 困難點

1. 要看出每個從起始點之前判斷式具有遞迴關係

## Solve Point

- [x]  建立一個 dp 為長度 len(s)+1 的boolean 矩陣用來儲存已運算過的結果，初始化最 dp[len(s)] = true 代表走到最後, 其他 dp[i] = false ，i = 0~ len(s) - 1 因為都還未比對
- [x]  逐步從 start = sLen - 1 到 start = 0, 對所有 符合 $start + len(word_k) ≤ sLen$ 且 $s[start: start+ len(word_k)] == word_k$ 找尋 更新 $dp[start] = dp[start+len(word_k)]$