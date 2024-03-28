package main

func longestCommonPrefix(strs []string) string {
	minLen := 100
	for _, str := range strs {
		if minLen > len(str) {
			minLen = len(str)
		}
	}

	var result string
Loop:
	for i := 0; i < minLen; i++ {
		var prefix string
		for j, str := range strs {
			c := str[i : i+1]
			if j == 0 {
				prefix = c
				continue
			}
			if prefix != c {
				prefix = ""
				break Loop
			}

		}

		result += prefix
	}

	return result
}

func longestCommonPrefix_2(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		strNum := len(strs[i])
		for j := 0; i < len(prefix); j++ {
			if strNum <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}
