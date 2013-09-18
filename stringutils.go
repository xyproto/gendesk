package main

import (
	"strings"
)

// Capitalize a string or return the same if it is too short
func capitalize(s string) string {
	if len(s) >= 2 {
		return strings.ToTitle(s[0:1]) + s[1:]
	}
	return s
}

// Checks if a trimmed line starts with a specific word
func startsWith(line string, word string) bool {
	//return 0 == strings.Index(strings.TrimSpace(line), word)
	return strings.HasPrefix(line, word)
}

// Return what's between two strings, "a" and "b", in another string
func between(orig string, a string, b string) string {
	if strings.Contains(orig, a) && strings.Contains(orig, b) {
		posa := strings.Index(orig, a) + len(a)
		posb := strings.LastIndex(orig, b)
		return orig[posa:posb]
	}
	return ""
}

// Return the contents between "" or '' (or an empty string)
func betweenQuotes(orig string) string {
	var s string
	for _, quote := range []string{"\"", "'"} {
		s = between(orig, quote, quote)
		if s != "" {
			return s
		}
	}
	return ""
}

// Return the string between the quotes or after the "=", if possible
// or return the original string
func betweenQuotesOrAfterEquals(orig string) string {
	s := betweenQuotes(orig)
	// Check for exactly one "="
	if (s == "") && (strings.Count(orig, "=") == 1) {
		s = strings.TrimSpace(strings.Split(orig, "=")[1])
	}
	return s
}

// Does a keyword exist in the string?
// Disregards several common special characters (like -, _ and .)
func has(s string, kw string) bool {
	lowercase := strings.ToLower(s)
	// Remove the most common special characters
	massaged := strings.Trim(lowercase, "-_.,!?()[]{}\\/:;+@")
	words := strings.Split(massaged, " ")
	for _, word := range words {
		if word == kw {
			return true
		}
	}
	return false
}

// Use a function for each element in a string list and
// return the modified list
func stringMap(f func(string) string, stringlist []string) []string {
	newlist := make([]string, len(stringlist))
	for i, elem := range stringlist {
		newlist[i] = f(elem)
	}
	return newlist
}
