package lib

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// FORMAT PACE
func FormatPace(p float64) string {
	s := fmt.Sprintf("%.2f", p)
	nums := strings.Split(s, ".")
	return fmt.Sprintf("%v'%v", nums[0], nums[1])
}

// FORMAT DURATION
func FormatDuration(minutes float64) string {
	mins := int(minutes)
	if mins < 60 {
		return fmt.Sprintf("%vm", mins)
	} else if mins < 1440 {
		h := mins / 60
		m := mins % 60
		if m == 0 {
			return fmt.Sprintf("%vh", h)
		}
		return fmt.Sprintf("%vh%vm", h, m)
	} else {
		d := mins / 1440
		rd := mins % 1440
		h := rd / 60
		m := rd % 60
		if m == 0 && h == 0 {
			return fmt.Sprintf("%vd", d)
		} else if m == 0 {
			return fmt.Sprintf("%vd%vh", d, h)
		} else if h == 0 {
			return fmt.Sprintf("%vd%vm", d, m)
		}
		return fmt.Sprintf("%vd%vh%vm", d, h, m)
	}
}

// FORMAT PERCENT
func FormatPercent(n float64) string {
	return fmt.Sprintf("%v%%", math.Round(n))
}

// PARSE STREAK
func ParseStreak(dates []string) int {
	longest, streak := 0, 1
	prev, _ := time.Parse(time.RFC822, dates[0]+" 10:00 MST")

	for i := 1; i < len(dates); i++ {
		current, _ := time.Parse(time.RFC822, dates[i]+" 10:00 MST")

		if prev.Add(24*time.Hour).Equal(current) || prev.Equal(current) {
			streak++
		} else {
			streak = 1
		}

		if longest < streak {
			longest = streak
		}

		prev = current
	}

	return longest
}

// PARSE RANGE
func ParseRange(dates []string) int {
	first, _ := time.Parse(time.RFC822, dates[0]+" 10:00 MST")
	last, _ := time.Parse(time.RFC822, dates[len(dates)-1]+" 10:00 MST")
	return int(last.Sub(first).Hours() / 24)
}
