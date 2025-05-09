package stringalign

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringAlignTestCase struct {
	input         string
	limit         int
	alignedString string
	aligner       func(str string, limit int) (string, error)
}

func TestStringAlign(t *testing.T) {
	tests := []stringAlignTestCase{
		{
			input:         "hello",
			limit:         10,
			alignedString: "hello     ",
			aligner:       LeftAlign,
		},
		{
			input:         "Serverless 🚀 computing is the future",
			limit:         15,
			alignedString: "Serverless 🚀  \ncomputing is   \nthe future     ",
			aligner:       LeftAlign,
		},
		{
			input:         "Hello, 🌍!\nThis is a long line that will be wrapped and aligned.",
			limit:         30,
			alignedString: "Hello, 🌍!                    \nThis is a long line that will \nbe wrapped and aligned.       ",
			aligner:       LeftAlign,
		},
		{
			input:         "hello",
			limit:         10,
			alignedString: "     hello",
			aligner:       RightAlign,
		},
		{
			input:         "lorem ipsum dolor sit amet",
			limit:         12,
			alignedString: " lorem ipsum\n   dolor sit\n        amet",
			aligner:       RightAlign,
		},
		{
			input:         "managing concurrent workers effectively in go routines",
			limit:         20,
			alignedString: " managing concurrent\n workers effectively\n      in go routines",
			aligner:       RightAlign,
		},
		{
			input:         "hello",
			limit:         11,
			alignedString: "   hello   ",
			aligner:       CenterAlign,
		},
		{
			input:         "Distributed systems are complex",
			limit:         16,
			alignedString: "  Distributed   \n  systems are   \n    complex     ",
			aligner:       CenterAlign,
		},
		{
			input:         "graceful shutdown and error propagation",
			limit:         24,
			alignedString: " graceful shutdown and  \n   error propagation    ",
			aligner:       CenterAlign,
		},
		{
			input:         "hello",
			limit:         15,
			alignedString: "hello",
			aligner:       Justify,
		},
		{
			input:         "The quick brown fox jumps over the lazy dog",
			limit:         12,
			alignedString: "The    quick\nbrown    fox\njumps   over\nthe lazy dog",
			aligner:       Justify,
		},
		{
			input:         "Implement robust logging and monitoring",
			limit:         20,
			alignedString: "Implement     robust\nlogging          and\nmonitoring",
			aligner:       Justify,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("Aligned String Test %d", idx+1), func(t *testing.T) {
			alignedString, _ := tt.aligner(tt.input, tt.limit)
			fmt.Println(alignedString)
			assert.Equal(t, tt.alignedString, alignedString)
		})
	}
}
