package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/warrenb95/cron-parser/internal/parser"
)

func TestParseFields(t *testing.T) {
	tests := map[string]struct {
		input string

		errContains string
		expected    parser.Fields
	}{
		"error zeroed out fields": {
			input:       "0 0 0 0 0 command",
			errContains: "value out of range",
		},
		"star expansion on all and tests limited response": {
			input: "* * * * * command",
			expected: parser.Fields{
				Minutes: func() []int { // TODO: could refactor to standalone func/constructor
					nums := make([]int, 60)
					for i := 0; i < 60; i++ {
						nums[i] = i
					}
					return nums
				}(),
				Hours: func() []int { // TODO: could refactor to standalone func/constructor
					nums := make([]int, 24)
					for i := 0; i < 24; i++ {
						nums[i] = i
					}
					return nums
				}(),
				DaysOfMonth: func() []int { // TODO: could refactor to standalone func/constructor
					nums := make([]int, 31)
					for i := 0; i < 31; i++ {
						nums[i] = i + 1
					}
					return nums
				}(),
				Months:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DaysOfWeek: []int{1, 2, 3, 4, 5, 6, 7},
				Command:    "command",
			},
		},
		"range expansion": {
			input: "1-5 1 1 1 1 command",
			expected: parser.Fields{
				Minutes:     []int{1, 2, 3, 4, 5},
				Hours:       []int{1},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{1},
				Command:     "command",
			},
		},
		"comma expansion": {
			input: "1,3,7 1 1 1 1 command",
			expected: parser.Fields{
				Minutes:     []int{1, 3, 7},
				Hours:       []int{1},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{1},
				Command:     "command",
			},
		},
		"step expansion with star": {
			input: "*/15 1 1 1 1 command",
			expected: parser.Fields{
				Minutes:     []int{0, 15, 30, 45},
				Hours:       []int{1},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{1},
				Command:     "command",
			},
		},
		"step expansion with range": {
			input: "1-10/2 1 1 1 1 command",
			expected: parser.Fields{
				Minutes:     []int{1, 3, 5, 7, 9},
				Hours:       []int{1},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{1},
				Command:     "command",
			},
		},
		"multiple tokens": {
			input: "1,2,5-7,*/20 1 1 1 1 command",
			expected: parser.Fields{
				Minutes:     []int{0, 1, 2, 5, 6, 7, 20, 40},
				Hours:       []int{1},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{1},
				Command:     "command",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fields, err := parser.ParseFields(test.input)
			if test.errContains != "" {
				require.ErrorContains(t, err, test.errContains, "error contains")
				return
			}
			require.NoError(t, err, "parsing fields")

			assert.Equal(t, test.expected, fields, "parsed fields")
		})
	}
}
