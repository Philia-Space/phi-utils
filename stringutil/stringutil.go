// Package stringutil provides common string manipulation helpers.
package stringutil

import (
	"strings"
	"unicode"
)

// Truncate truncates a string to maxLength, adding ellipsis if truncated.
func Truncate(s string, maxLength int) string {
	if maxLength <= 0 {
		return ""
	}
	if len(s) <= maxLength {
		return s
	}
	if maxLength <= 3 {
		return s[:maxLength]
	}
	return s[:maxLength-3] + "..."
}

// TruncateWords truncates a string to maxWords words, adding ellipsis if truncated.
func TruncateWords(s string, maxWords int) string {
	words := strings.Fields(s)
	if len(words) <= maxWords {
		return s
	}
	return strings.Join(words[:maxWords], " ") + "..."
}

// Slugify converts a string to a URL-friendly slug.
func Slugify(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			result.WriteRune(unicode.ToLower(r))
		} else if r == ' ' || r == '-' || r == '_' {
			result.WriteRune('-')
		}
	}
	return strings.Trim(result.String(), "-")
}

// CamelToSnake converts CamelCase to snake_case.
func CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// SnakeToCamel converts snake_case to CamelCase.
func SnakeToCamel(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// KebabToCamel converts kebab-case to CamelCase.
func KebabToCamel(s string) string {
	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// ContainsAny reports whether s contains any of the substrings.
func ContainsAny(s string, substrs ...string) bool {
	for _, sub := range substrs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// RemoveDuplicates removes duplicate strings from a slice, preserving order.
func RemoveDuplicates(ss []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(ss))
	for _, s := range ss {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}

// Intersection returns strings present in both slices.
func Intersection(a, b []string) []string {
	seen := make(map[string]bool)
	for _, s := range a {
		seen[s] = true
	}
	var result []string
	for _, s := range b {
		if seen[s] {
			result = append(result, s)
			delete(seen, s)
		}
	}
	return result
}

// IsBlank reports whether a string is empty or whitespace only.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank reports whether a string has non-whitespace content.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// Coalesce returns the first non-blank string, or empty if all are blank.
func Coalesce(values ...string) string {
	for _, v := range values {
		if IsNotBlank(v) {
			return v
		}
	}
	return ""
}

// PadLeft pads a string with a rune to reach target length.
func PadLeft(s string, length int, pad rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(pad), length-len(s))
	return padding + s
}

// PadRight pads a string with a rune to reach target length.
func PadRight(s string, length int, pad rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(pad), length-len(s))
	return s + padding
}

// Reverse reverses a string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// FirstN returns the first n runes of a string.
func FirstN(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n])
}

// LastN returns the last n runes of a string.
func LastN(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[len(runes)-n:])
}
