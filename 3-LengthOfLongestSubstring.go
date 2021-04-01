package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	var longestSubstring string
	var currSubstring string

	for _, val := range s {
		currChar := string(val)

		if !strings.Contains(currSubstring, currChar) {
			currSubstring = currSubstring + currChar
		} else {
			if len(currSubstring) > len(longestSubstring) {
				longestSubstring = currSubstring
			}
			currSubstring = currSubstring[strings.Index(currSubstring, currChar)+1:] + currChar
		}
	}

	if len(currSubstring) > len(longestSubstring) {
		longestSubstring = currSubstring
	}

	return len(longestSubstring)
}
