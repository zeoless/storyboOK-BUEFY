package source

import (
	"fmt"
	"testing"
)

// newEdit is a testing helper for quickly generating Edits
func newEdit(loc int, old, new string) Edit {
	return Edit{Location: loc, Old: old, New: new}
}

// TestMutateSingle tests almost every possibility of a single edit
func TestMutateSingle(t *testing.T) {
	type test struct {
		input    string
		edit     Edit
		expected string
	}

	tests := []test{
		// Simple edits that replace everything
		{"", newEdit(0, "", ""), ""},
		{"a", newEdit(0, "a", "A"), "A"},
		{"abcde", newEdit(0, "abcde", "fghij"), "fghij"},
		{"", newEdit(0, "", "fghij"), "fghij"},
		{"abcde", newEdit(0, "abcde", ""), ""},

		// Edits that start at the very beginning (But don't cover the whole range)
		{"abcde", newEdit(0, "a", "A"), "Abcde"},
		{"abcde", newEdit(0, "ab", "AB"), "ABcde"},
		{"abcde", newEdit(0, "abc", "ABC"), "ABCde"},
		{"abcde", newEdit(0, "abcd", "ABCD"), "ABCDe"},

		// The above repeated, but with different lengths
		{"abcde", newEdit(0, "a", ""), "bcde"},
		{"abcde", newEdit(0, "ab", "A"), "Acde"},
		{"abcde", newEdit(0, "abc", "AB"), "ABde"},
		{"abcde", newEdit(0, "abcd", "AB"), "ABe"},

		// Edits that touch the end (but don't cover the whole range)
		{"abcde", newEdit(4, "e", "E"), "abcdE"},
		{"abcde", newEdit(3, "de", "DE"), "abcDE"},
		{"abcde", newEdit(2, "cde", "CDE"), "abCDE"},
		{"abcde", newEdit(1, "bcde", "BCDE"), "aBCDE"},

		// The above repeated, but with different lengths
		{"abcde", newEdit(4, "e", ""), "abcd"},
		{"abcde", newEdit(3, "de", "D"), "abcD"},
		{"abcde", newEdit(2, "cde", "CD"), "abCD"},
		{"abcde", newEdit(1, "bcde", "BC"), "aBC"},

		// Raw insertions / deletions
		{"abcde", newEdit(0, "", "_"), "_abcde"},
		{"abcde", newEdit(1, "", "_"), "a_bcde"}