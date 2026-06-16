package scheduler

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	format = "20060102"
)

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func nextYear(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(format, dstart)
	if err != nil {
		return "", err
	}

	for {
		date = date.AddDate(1, 0, 0)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format(format), nil
}

func nextDay(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(format, dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Split(repeat, " ")
	if len(parts) != 2 {
		return "", errors.New("the repeat parameter is invalid")
	}

	interval, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	if interval < 1 || interval > 400 {
		return "", errors.New("repeat: invalid number")
	}

	for {
		date = date.AddDate(0, 0, interval)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format(format), nil
}

func nextWeek(now time.Time, dstart string, repeat string) (string, error) {

	date, err := time.Parse(format, dstart)
	if err != nil {
		return "", err
	}

	var week [7]bool
	parts := strings.Split(repeat, " ")
	if len(parts) != 2 {
		return "", errors.New("the repeat parameter is invalid")
	}

	weekDays := strings.Split(parts[1], ",")
	for _, value := range weekDays {
		num, err := strconv.Atoi(value)
		if err != nil {
			return "", err
		}

		if num < 1 || num > 7 {
			return "", errors.New("repeat: invalid number")
		}

		if num == 7 {
			num = 0
		}

		week[num] = true
	}

	for {
		date = date.AddDate(0, 0, 1)
		if week[int(date.Weekday())] && afterNow(date, now) {
			break
		}
	}

	return date.Format(format), nil
}

func nextMonth(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(format, dstart)
	if err != nil {
		return "", err
	}

	var day [32]bool
	var month [13]bool
	var months []string
	allowLastDay := false
	allowBeforeLastDay := false

	parts := strings.Split(repeat, " ")
	if len(parts) < 2 || len(parts) > 3 {
		return "", errors.New("the repeat parameter is invalid")
	}

	monthDays := strings.Split(parts[1], ",")

	for _, value := range monthDays {
		num, err := strconv.Atoi(value)
		if err != nil {
			return "", err
		}

		if num == -1 {
			allowLastDay = true
			continue
		}
		if num == -2 {
			allowBeforeLastDay = true
			continue
		}

		if num < 1 || num > 31 {
			return "", errors.New("repeat: invalid number")
		}

		day[num] = true
	}

	if len(parts) == 3 {
		months = strings.Split(parts[2], ",")

		for _, value := range months {
			num, err := strconv.Atoi(value)
			if err != nil {
				return "", err
			}

			if num < 1 || num > 12 {
				return "", errors.New("repeat: invalid number")
			}

			month[num] = true
		}
	}

	for {
		matched := false
		monthMatched := false

		date = date.AddDate(0, 0, 1)

		if len(parts) == 3 {
			monthMatched = month[int(date.Month())]
		} else {
			monthMatched = true
		}

		if day[date.Day()] {
			matched = true
		}
		if allowLastDay && isLastDay(date) {
			matched = true
		}
		if allowBeforeLastDay && isBeforeLastDay(date) {
			matched = true
		}

		if matched && monthMatched && afterNow(date, now) {
			break
		}
	}

	return date.Format(format), nil
}

func isLastDay(date time.Time) bool {
	month := date.Month()

	date = date.AddDate(0, 0, 1)
	nextMonth := date.Month()

	if month != nextMonth {
		return true
	}

	return false
}

func isBeforeLastDay(date time.Time) bool {
	month := date.Month()

	date = date.AddDate(0, 0, 2)
	nextMonth := date.Month()

	if month != nextMonth {
		return true
	}

	return false
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("the repeat parameter is empty")
	}

	switch {
	case repeat == "y":
		return nextYear(now, dstart, repeat)

	case strings.HasPrefix(repeat, "d"):
		return nextDay(now, dstart, repeat)

	case strings.HasPrefix(repeat, "w"):
		return nextWeek(now, dstart, repeat)

	case strings.HasPrefix(repeat, "m"):
		return nextMonth(now, dstart, repeat)

	default:
		return "", errors.New("invalid rule")
	}
}
