package str

import "strings"

// 连接字符串
func Join(strs ...string) string {
	var builder strings.Builder
	if len(strs) == 0 {
		return ""
	}
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}
