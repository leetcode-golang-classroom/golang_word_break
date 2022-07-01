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
			}
			if dp[start] {
				break
			}
		}
	}
	return dp[0]
}
