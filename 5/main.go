package _5

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true

		if i < len(s)-1 {
			dp[i][i+1] = s[i] == s[i+1]
		}
	}

	for i := 2; i <= len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			dp[j][j+i] = s[j] == s[j+i] && dp[j+1][j+i-1]
		}
	}

	result := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if dp[i][j] && len(s[i:j+1]) > len(result) {
				result = s[i : j+1]
			}
		}
	}

	return result
}

func isPalindrome(s string) bool {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return false
}
