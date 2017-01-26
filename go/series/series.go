package slice

func All(n int, s string) []string {
	l := len(s)
	res := make([]string, l-n+1)
	for i := 0; i < l-n+1; i++ {
		res[i] = s[i : i+n]
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		return s[:n], true
	}
	return "", false
}
