package levenshtein 

const r rune = 0

func Distance(s1 string, s2 string) int {
	rs1 := append([]rune{r}, []rune(s1)...)
	rs2 := append([]rune{r}, []rune(s2)...)

	dp := make([][]int, len(rs1))
	for i := 0; i < len(rs1); i++ {
		dp[i] = make([]int, len(rs2))
		for j := 0; j < len(rs2); j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			} else if j == 0 {
				dp[i][j] = i
				continue
			}

			s := dp[i-1][j-1]		

			if rs1[i] != rs1[i-1] {
		  	s = min(s, dp[i-1][j])
		  }

		  if rs2[j] != rs2[j-1] {
		  	s = min(s, dp[i][j-1])
		  }

			if rs1[i] == rs2[j] {
				dp[i][j] = s 
			} else {
				dp[i][j] = s + 1
			}
		}
	}

	return dp[len(rs1)-1][len(rs2)-1]
}
