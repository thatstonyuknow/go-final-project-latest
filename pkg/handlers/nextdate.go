package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// NextDate returns the next date for task repetition
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Check for empty repeat string
	if strings.TrimSpace(repeat) == "" {
		return "", fmt.Errorf("repeat rule cannot be empty")
	}

	// Parse the start date
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", fmt.Errorf("invalid start date format: %v", err)
	}

	// Parse the repeat rule
	parts := strings.Split(strings.TrimSpace(repeat), " ")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid repeat format")
	}

	rule := parts[0]

	switch rule {
	case "d":
		return handleDayRule(date, now, parts)
	case "y":
		return handleYearRule(date, now)
	case "w":
		return handleWeekRule(date, now, parts)
	case "m":
		return handleMonthRule(date, now, parts)
	default:
		return "", fmt.Errorf("unsupported repeat format: %s", rule)
	}
}

// afterNow checks if the first date is after the second (considering only date, ignoring time)
func afterNow(date, now time.Time) bool {
	dateOnly := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return dateOnly.After(nowOnly)
}

// handleDayRule handles the "d <number>" rule
func handleDayRule(date, now time.Time, parts []string) (string, error) {
	if len(parts) != 2 {
		return "", fmt.Errorf("day rule requires interval")
	}

	interval, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid day interval: %v", err)
	}

	if interval <= 0 || interval > 400 {
		return "", fmt.Errorf("day interval must be between 1 and 400")
	}

	for {
		date = date.AddDate(0, 0, interval)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format("20060102"), nil
}

// handleYearRule handles the "y" rule
func handleYearRule(date, now time.Time) (string, error) {
	for {
		date = date.AddDate(1, 0, 0)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format("20060102"), nil
}

// handleWeekRule handles the "w <weekdays>" rule
func handleWeekRule(date, now time.Time, parts []string) (string, error) {
	if len(parts) != 2 {
		return "", fmt.Errorf("week rule requires days")
	}

	daysStr := strings.Split(parts[1], ",")
	validDays := make(map[int]bool)

	for _, dayStr := range daysStr {
		day, err := strconv.Atoi(strings.TrimSpace(dayStr))
		if err != nil {
			return "", fmt.Errorf("invalid day of week: %v", err)
		}
		if day < 1 || day > 7 {
			return "", fmt.Errorf("day of week must be between 1 and 7")
		}
		validDays[day] = true
	}

	// Start from the date after now and find the first matching day
	current := date
	if !afterNow(current, now) {
		current = now.AddDate(0, 0, 1)
	}

	for {
		// Convert Go weekday (0=Sunday) to our format (1=Monday, 7=Sunday)
		weekday := int(current.Weekday())
		if weekday == 0 {
			weekday = 7
		}

		if validDays[weekday] {
			if afterNow(current, now) {
				return current.Format("20060102"), nil
			}
		}
		current = current.AddDate(0, 0, 1)
	}
}

// handleMonthRule handles the "m <month_days> [months]" rule
func handleMonthRule(date, now time.Time, parts []string) (string, error) {
	if len(parts) < 2 {
		return "", fmt.Errorf("month rule requires days")
	}

	daysStr := strings.Split(parts[1], ",")
	validDays := make(map[int]bool)

	for _, dayStr := range daysStr {
		day, err := strconv.Atoi(strings.TrimSpace(dayStr))
		if err != nil {
			return "", fmt.Errorf("invalid day of month: %v", err)
		}
		if day < -2 || day == 0 || day > 31 {
			return "", fmt.Errorf("day of month must be between 1-31 or -1, -2")
		}
		validDays[day] = true
	}

	validMonths := make(map[int]bool)
	if len(parts) >= 3 {
		monthsStr := strings.Split(parts[2], ",")
		for _, monthStr := range monthsStr {
			month, err := strconv.Atoi(strings.TrimSpace(monthStr))
			if err != nil {
				return "", fmt.Errorf("invalid month: %v", err)
			}
			if month < 1 || month > 12 {
				return "", fmt.Errorf("month must be between 1 and 12")
			}
			validMonths[month] = true
		}
	} else {
		// If months are not specified, use all months
		for i := 1; i <= 12; i++ {
			validMonths[i] = true
		}
	}

	// Start from the date after now
	current := date
	if !afterNow(current, now) {
		current = now.AddDate(0, 0, 1)
	}

	for {
		month := int(current.Month())
		if validMonths[month] {
			day := current.Day()
			lastDay := getLastDayOfMonth(current.Year(), current.Month())

			// Check regular days
			if validDays[day] {
				if afterNow(current, now) {
					return current.Format("20060102"), nil
				}
			}

			// Check negative days (-1 = last day, -2 = second to last)
			if validDays[-1] && day == lastDay {
				if afterNow(current, now) {
					return current.Format("20060102"), nil
				}
			}
			if validDays[-2] && day == lastDay-1 {
				if afterNow(current, now) {
					return current.Format("20060102"), nil
				}
			}
		}
		current = current.AddDate(0, 0, 1)
	}
}

// getLastDayOfMonth returns the last day of the month
func getLastDayOfMonth(year int, month time.Month) int {
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastDay := nextMonth.AddDate(0, 0, -1)
	return lastDay.Day()
}