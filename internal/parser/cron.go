package parser

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Fields is a structured format for cron inputs.
type Fields struct {
	Minutes     []int
	Hours       []int
	DaysOfMonth []int
	Months      []int
	DaysOfWeek  []int
	Command     string
}

// ParseFields will parse the input cron string and format it into a
// Fields struct that can then more easily be used.
func ParseFields(input string) (Fields, error) {
	parts := strings.Fields(strings.TrimSpace(input))
	if len(parts) != 6 {
		return Fields{}, errors.New("invalid cron expression")
	}

	var f Fields
	var err error

	f.Minutes, err = expandField(parts[0], 0, 59)
	if err != nil {
		return Fields{}, err
	}

	f.Hours, err = expandField(parts[1], 0, 23)
	if err != nil {
		return Fields{}, err
	}

	f.DaysOfMonth, err = expandField(parts[2], 1, 31)
	if err != nil {
		return Fields{}, err
	}

	f.Months, err = expandField(parts[3], 1, 12)
	if err != nil {
		return Fields{}, err
	}

	f.DaysOfWeek, err = expandField(parts[4], 1, 7)
	if err != nil {
		return Fields{}, err
	}

	f.Command = parts[5]
	return f, nil
}

// expandField takes in the field token and min max values
// calculates and then returns the expected values.
func expandField(field string, min, max int) ([]int, error) {
	tokens := strings.Split(field, ",")
	resultMap := make(map[int]struct{})

	for _, token := range tokens {
		token = strings.TrimSpace(token)

		var values []int
		var err error

		if strings.Contains(token, "/") {
			values, err = expandStep(token, min, max)
			if err != nil {
				return nil, err
			}
		} else if token == "*" {
			values = rangeList(min, max, 1)
		} else if strings.Contains(token, "-") {
			values, err = expandRange(token, min, max)
			if err != nil {
				return nil, err
			}
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}

			if num < min || num > max {
				return nil, errors.New("value out of range")
			}

			values = []int{num}
		}

		for _, v := range values {
			resultMap[v] = struct{}{}
		}
	}

	var result []int
	for k := range resultMap {
		result = append(result, k)
	}

	sort.Ints(result)

	return result, nil
}

// expandStep will calculate and return all the values in the step token.
func expandStep(token string, min, max int) ([]int, error) {
	parts := strings.Split(token, "/")
	if len(parts) != 2 {
		return nil, errors.New("invalid step expression")
	}

	step, err := strconv.Atoi(parts[1])
	if err != nil || step <= 0 {
		return nil, errors.New("invalid step value")
	}

	var baseValues []int
	if parts[0] == "*" {
		baseValues = rangeList(min, max, 1)
	} else if strings.Contains(parts[0], "-") {
		baseValues, err = expandRange(parts[0], min, max)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid base for step expression")
	}

	var stepped []int
	for i := 0; i < len(baseValues); i += step {
		stepped = append(stepped, baseValues[i])
	}
	return stepped, nil
}

// expandRange will expand on a range, returning all values within the range.
func expandRange(token string, min, max int) ([]int, error) {
	parts := strings.Split(token, "-")
	if len(parts) != 2 {
		return nil, errors.New("invalid range expression")
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	if start > end || start < min || end > max {
		return nil, fmt.Errorf("range out of bounds, min: %d, max: %d", min, max)
	}
	return rangeList(start, end, 1), nil
}

// rangeList is simple function to create a []int.
func rangeList(start, end, step int) []int {
	var result []int
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result
}
