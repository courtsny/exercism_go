package main

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(log_line string) bool {
	log_chunks := strings.Split(log_line, " ")
	// I know this is prob supposed to be a regex, but this seems cleaner to me tbh
	switch log_chunks[0] {
	case "[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]":
		return true
	}
	return false
}

func SplitLogLine(log_line string) []string {
	// In between < >, any of ~, * (needs esc char), =, - can be in there 1+ time(s)
	re := regexp.MustCompile(`<[~\*=-]+>`)
	log_chunks := re.Split(log_line, -1)
	//fmt.Println(log_chunks)
	return log_chunks
}

func CountQuotedPasswords(log_array []string) int {
	// Ignore casing, find password anywhere in the string
	re := regexp.MustCompile(`(?i)".*password.*"`)
	log_chunks := 0
	for _, x := range log_array {
		match := re.FindString(x)
		if len(match) > 0 {
			//fmt.Printf("found %s in %s\n", match, x)
			log_chunks += 1
		}
	}
	return log_chunks
}

func RemoveEndOfLineText(log_line string) string {
	// Take out end-of-line + whatever digits come after
	re := regexp.MustCompile(`end-of-line\d*`)
	log_chunks := re.Split(log_line, -1)
	cleaned_log := strings.Join(log_chunks, "")
	return cleaned_log
}

func TagWithUserName(log_array []string) []string {
	re := regexp.MustCompile(`User\s+[a-zA-Z\d]*`)
	username_logs := []string{}
	for _, x := range log_array {
		match := re.FindString(x)
		if len(match) > 0 {
			// Get User + Name, and grab the Name
			username_array := strings.Split(match, " ")
			username := username_array[len(username_array)-1]
			// Remove User + Name and rejoin the string
			chunked_log := re.Split(x, -1)
			cleaned_log := strings.Join(chunked_log, "")
			// Add [USR] and Name to the front of the cleaned string
			cleaned_log = "[USR] " + username + " " + cleaned_log
			username_logs = append(username_logs, cleaned_log)
			//fmt.Println("cleaned log:", cleaned_log)
		} else {
			// If there's nothing to clean, add the log as is
			username_logs = append(username_logs, x)
		}
	}
	return username_logs
}

func PrintLogArray(log_array []string) string {
	var sb strings.Builder
	for _, x := range log_array {
		sb.WriteString(x + "\n")
	}
	return sb.String()
}

func RegexesTest() {
	fmt.Printf("example logs:\nexpect true: %t\nexpect false: %t, %t\n", IsValidLine("[ERR] A good error here"), IsValidLine("Any old [ERR] text"), IsValidLine("[BOB] Any old text"))
	arrow_log := "section 1<*>section 2<~~~>section 3"
	chunked_arrow_log := SplitLogLine(arrow_log)
	fmt.Printf("arrow log before split: %s\narrow log after split: %s\n\n", arrow_log, chunked_arrow_log)
	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	password_count := CountQuotedPasswords(lines)
	fmt.Printf("found %d passwords (case insensitive) in logs:\n %s\n", password_count, PrintLogArray(lines))
	end_of_line := "[INF] end-of-line23033 Network Failure end-of-line27"
	cleaned_end_of_line := RemoveEndOfLineText(end_of_line)
	fmt.Printf("before end-of-line removal: %s\nafter end-of-line removal: %s\n\n", end_of_line, cleaned_end_of_line)
	user_lines := []string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	}
	cleaned_user_lines := TagWithUserName(user_lines)
	fmt.Printf("raw logs: \n%s\ncleaned logs: \n%s\n", PrintLogArray(user_lines), PrintLogArray(cleaned_user_lines))
}
