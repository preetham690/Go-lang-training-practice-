package strings

func Reverse(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		ans = string(s[i]) + ans
	}
	return ans
}
