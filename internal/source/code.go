
package source

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Edit struct {
	Location int
	Old      string
	New      string
}

func LineNumber(source string, head int) (int, int) {
	// Calculate the true line and column number for a query, ignoring spaces
	var comment bool
	var loc, line, col int
	for i, char := range source {
		loc += 1
		col += 1
		// TODO: Check bounds
		if char == '-' && source[i+1] == '-' {
			comment = true
		}
		if char == '\n' {
			comment = false
			line += 1
			col = 0
		}
		if loc <= head {
			continue
		}
		if unicode.IsSpace(char) {
			continue
		}
		if comment {
			continue
		}
		break
	}
	return line + 1, col
}

func Pluck(source string, location, length int) (string, error) {
	head := location
	tail := location + length
	return source[head:tail], nil
}

func Mutate(raw string, a []Edit) (string, error) {
	if len(a) == 0 {
		return raw, nil
	}

	sort.Slice(a, func(i, j int) bool { return a[i].Location > a[j].Location })

	s := raw
	for idx, edit := range a {