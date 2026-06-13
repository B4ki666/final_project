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

	/*case strings.HasPrefix(repeat, "m"):
	return nextMonth(now, dstart, repeat)*/

	default:
		return "", errors.New("invalid rule")
	}
}
