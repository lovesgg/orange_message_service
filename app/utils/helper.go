package utils

func InStringArray(s string, ss []string) bool {
	for _, item := range ss {
		if item == s {
			return true
		}
	}
	return false
}

func GetStringByIndex(ss []string, index int) string {
	if index > len(ss) {
		return ""
	}
	return ss[index]
}
