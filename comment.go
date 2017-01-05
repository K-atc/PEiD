package main

import (
	// "fmt"
	"sort"
	// "strings"
	// "github.com/fatih/color"
)

func find_string(haystack []string, needle string) bool {
	i := sort.SearchStrings(haystack, needle)
	if i < len(haystack) && haystack[i] == needle {
		return true
	}
	return false
}

func make_comment_line(title, data string) string {
	// if strings.Contains(data, "OK") || strings.Contains(data, "ok") ||
	// strings.Contains(data, "YES") || strings.Contains(data, "yes") {
	//     data = color.Green(data)
	// } else if strings.Contains(data, "NG") || strings.Contains(data, "NG") ||
	// strings.Contains(data, "NO") || strings.Contains(data, "no") {
	//     data = color.Red(data)
	// }
	if len(data) == 0 {
		return "  " + title + "\n"
	} else {
		return "  " + title + " : " + data + "\n"
	}
}

func add_comment(record YaraRecord) (string, bool) {
	var comment string

	sorted_matched_rules := make([]string, len(record.matched_rules))
	copy(sorted_matched_rules, record.matched_rules)
	sort.Strings(sorted_matched_rules)

	var is_PE bool
	if find_string(sorted_matched_rules, "IsPE32") {
		comment += make_comment_line("PE", "32 bit")
		is_PE = true
	} else if find_string(sorted_matched_rules, "IsPE64") {
		comment += make_comment_line("PE", "64 bit")
		is_PE = true
	} else {
		is_PE = false
	}

	var is_packed bool
	if is_PE {
		if find_string(sorted_matched_rules, "IsDLL") {
			comment += make_comment_line("DLL", "yes")
		} else {
			comment += make_comment_line("DLL", "no")
		}
		if find_string(sorted_matched_rules, "anti_dbg") {
			comment += make_comment_line("Anti-Debug", "yes")
		} else {
			comment += make_comment_line("Anti-Debug", "no")
		}
		if find_string(sorted_matched_rules, "IsPacked") {
			comment += make_comment_line("Packed", "yes")
			is_packed = true
		} else {
			comment += make_comment_line("Packed", "no")
			is_packed = false
		}
		if find_string(sorted_matched_rules, "IsWindowsGUI") {
			comment += make_comment_line("GUI Program", "yes")
		} else {
			if is_packed {
				comment += make_comment_line("GUI Program", "no (yes)")
			} else {
				comment += make_comment_line("GUI Program", "no")
			}
		}
		if find_string(sorted_matched_rules, "IsConsole") {
			comment += make_comment_line("Console Program", "yes")
		} else {
			if is_packed {
				comment += make_comment_line("Console Program", "no (yes)")
			} else {
				comment += make_comment_line("Console Program", "no")
			}
		}
		if find_string(sorted_matched_rules, "win_mutex") {
			comment += make_comment_line("mutex", "yes")
		}
	}

	if find_string(sorted_matched_rules, "contentis_base64") {
		comment += make_comment_line("contains base64", "")
	}
	if find_string(sorted_matched_rules, "with_urls") {
		comment += make_comment_line("contains urls", "")
	}

	if len(comment) > 0 {
		return comment, true
	} else {
		return "", false
	}
}
