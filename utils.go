package main

import (
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Capitalize a string or return the same if it is too short
func capitalize(s string) string {
	switch {
	case len(s) >= 2:
		return strings.ToTitle(s[0:1]) + s[1:]
	case len(s) == 1:
		return strings.ToUpper(string(s[0]))
	default:
		return s
	}
}

// Return what's between two strings, "a" and "b", in another string
func between(orig string, a string, b string) string {
	if strings.Contains(orig, a) && strings.Contains(orig, b) {
		posa := strings.Index(orig, a) + len(a)
		posb := strings.LastIndex(orig, b)
		if posa > posb {
			return ""
		}
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
	if s == "" && strings.Count(orig, "=") == 1 {
		s = strings.TrimSpace(strings.Split(orig, "=")[1])
	}
	return s
}

// Does a keyword exist in the string?
// Disregards several common special characters (like -, _ and .)
func has(s string, kw string) bool {
	var (
		lowercase = strings.ToLower(s)
		// Remove the most common special characters
		massaged = strings.Trim(lowercase, "-_.,!?()[]{}\\/:;+@")
		words    = strings.Split(massaged, " ")
	)
	for _, word := range words {
		if word == kw {
			return true
		}
	}
	return false
}
