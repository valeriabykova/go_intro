func reverse(s string) string {
	copy := []rune(s)
	for i := 0; i < len(s); i++ {
		copy[len(s)-i-1] = s[i]
	}
	return string(copy)
}
