package dates

import (
	"fmt"
	"time"
)

func LongDateString(data time.Time) string {
	day := data.Day()

	if day == 11 || day == 12 || day == 13 {
		return fmt.Sprintf("the %dth of %s %d", day, data.Month(), data.Year())
	}

	firstDayDigit := day % 10

	var ending string
	switch firstDayDigit {
	case 1:
		ending = "st"
	case 2:
		ending = "nd"
	case 3:
		ending = "rd"
	default:
		ending = "th"
	}

	return fmt.Sprintf("the %d%s of %s %d", day, ending, data.Month(), data.Year())
}

func ShortDateString(data time.Time) string {
	return fmt.Sprintf("%d %s %d", data.Day(), data.Month().String()[:3], data.Year())
}
