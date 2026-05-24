// Package timeutil provides common time formatting and manipulation helpers.
package timeutil

import (
	"fmt"
	"time"
)

// ISO8601Layout is the ISO 8601 datetime format.
const ISO8601Layout = "2006-01-02T15:04:05Z"

// ISO8601WithMsLayout includes milliseconds.
const ISO8601WithMsLayout = "2006-01-02T15:04:05.000Z"

// DateLayout is the date-only format.
const DateLayout = "2006-01-02"

// ToISO8601 formats a time as ISO 8601 string.
func ToISO8601(t time.Time) string {
	return t.UTC().Format(ISO8601Layout)
}

// FromISO8601 parses an ISO 8601 string.
func FromISO8601(s string) (time.Time, error) {
	return time.Parse(ISO8601Layout, s)
}

// ToISO8601WithMs formats a time with millisecond precision.
func ToISO8601WithMs(t time.Time) string {
	return t.UTC().Format(ISO8601WithMsLayout)
}

// StartOfDay returns the start of the day for a given time.
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day for a given time.
func EndOfDay(t time.Time) time.Time {
	return StartOfDay(t).Add(24*time.Hour - time.Nanosecond)
}

// StartOfWeek returns the start of the week (Sunday) for a given time.
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	return StartOfDay(t.AddDate(0, 0, -weekday))
}

// EndOfWeek returns the end of the week (Saturday) for a given time.
func EndOfWeek(t time.Time) time.Time {
	return StartOfWeek(t).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// StartOfMonth returns the start of the month.
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the end of the month.
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// RelativeTime returns a human-readable relative time string.
func RelativeTime(t time.Time) string {
	duration := time.Since(t)
	if duration < 0 {
		duration = -duration
		return fmt.Sprintf("in %s", humanDuration(duration))
	}
	return fmt.Sprintf("%s ago", humanDuration(duration))
}

func humanDuration(d time.Duration) string {
	if d < time.Minute {
		return "a few seconds"
	}
	if d < time.Hour {
		mins := int(d.Minutes())
		if mins == 1 {
			return "a minute"
		}
		return fmt.Sprintf("%d minutes", mins)
	}
	if d < 24*time.Hour {
		hours := int(d.Hours())
		if hours == 1 {
			return "an hour"
		}
		return fmt.Sprintf("%d hours", hours)
	}
	if d < 7*24*time.Hour {
		days := int(d.Hours() / 24)
		if days == 1 {
			return "a day"
		}
		return fmt.Sprintf("%d days", days)
	}
	if d < 30*24*time.Hour {
		weeks := int(d.Hours() / 24 / 7)
		if weeks == 1 {
			return "a week"
		}
		return fmt.Sprintf("%d weeks", weeks)
	}
	if d < 365*24*time.Hour {
		months := int(d.Hours() / 24 / 30)
		if months == 1 {
			return "a month"
		}
		return fmt.Sprintf("%d months", months)
	}
	years := int(d.Hours() / 24 / 365)
	if years == 1 {
		return "a year"
	}
	return fmt.Sprintf("%d years", years)
}

// MustParse parses a time string or panics.
func MustParse(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(fmt.Sprintf("timeutil: failed to parse %q with layout %q: %v", value, layout, err))
	}
	return t
}

// TruncateToMinute truncates a time to minute precision.
func TruncateToMinute(t time.Time) time.Time {
	return t.Truncate(time.Minute)
}

// IsSameDay reports whether two times are on the same calendar day.
func IsSameDay(a, b time.Time) bool {
	y1, m1, d1 := a.Date()
	y2, m2, d2 := b.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// AgeInYears calculates age in years from a birth date.
func AgeInYears(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}
