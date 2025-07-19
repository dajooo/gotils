package strutil

import "strings"

func JoinStream(sep string, stringChan <-chan string) string {
	var builder strings.Builder

	first, ok := <-stringChan
	if !ok {
		return ""
	}
	builder.WriteString(first)

	for s := range stringChan {
		builder.WriteString(sep)
		builder.WriteString(s)
	}

	return builder.String()
}
