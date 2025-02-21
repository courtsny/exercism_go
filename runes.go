package main

import (
	"fmt"
	"strings"
)

func Application(log_line string) string {
	for _, x := range log_line {
		if x == rune('❗') {
			return "recommendation"
		}
		if x == rune('🔍') {
			return "search"
		}
		if x == rune('☀') {
			return "weather"
		}
	}
	return ""
}

func Replace(log_line string, bad_char, new_char rune) string {
	var sb strings.Builder
	for _, x := range log_line {
		if x == bad_char {
			sb.WriteRune(new_char)
		} else {
			sb.WriteRune(x)
		}
	}
	return sb.String()
}

func WithinLimit(log_line string, limit int) bool {
	length := 0
	for length = range log_line {
		length += 1
	}
	if length <= limit {
		return true
	}
	return false
}

func RunesTest() {
	log_string := "❗ recommended search product 🔍"
	fmt.Printf("log string '%s', first application is '%s'\n", log_string, Application(log_string))
	log := "please replace '👎' with '👍'"
	new_log := Replace(log, '👎', '👍')
	fmt.Printf("original log: %s\nnew log: %s\n", log, new_log)
	hello_log := "hello❗"
	fmt.Printf("'%s' is within the limit: %t\n", hello_log, WithinLimit(hello_log, 6))
}
