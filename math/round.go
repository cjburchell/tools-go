package math

import (
	"fmt"
	"strconv"
)

// Round the value down to given digits
func Round(value float64, digits int) float64 {
	format := fmt.Sprintf("%%.%df", digits)
	formatted, _ := strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return formatted
}
