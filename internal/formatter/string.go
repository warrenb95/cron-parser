package formatter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/warrenb95/cron-parser/internal/parser"
)

// Format builds a human-readable string from the parsed cron fields.
// If a field contains more than 14 values, only the first 14 are shown.
func Format(fields parser.Fields) string {
	var builder strings.Builder

	// TODO: make the 14 configureable?
	builder.WriteString(fmt.Sprintf("%-14s%s\n", "minute", formatSlice(fields.Minutes)))
	builder.WriteString(fmt.Sprintf("%-14s%s\n", "hour", formatSlice(fields.Hours)))
	builder.WriteString(fmt.Sprintf("%-14s%s\n", "day of month", formatSlice(fields.DaysOfMonth)))
	builder.WriteString(fmt.Sprintf("%-14s%s\n", "month", formatSlice(fields.Months)))
	builder.WriteString(fmt.Sprintf("%-14s%s\n", "day of week", formatSlice(fields.DaysOfWeek)))
	builder.WriteString(fmt.Sprintf("%-14s%s", "command", fields.Command))

	return builder.String()
}

func formatSlice(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	// TODO: make configureable?
	if len(nums) > 14 {
		nums = nums[:14]
	}

	var strSlice []string
	for _, n := range nums {
		strSlice = append(strSlice, strconv.Itoa(n))
	}
	return strings.Join(strSlice, " ")
}
