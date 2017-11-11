package calendar

import (
	"fmt"
	"strings"
	"time"
)

type Calendar struct {
	WeekLabels   []string
	MonthLabels  []string
	Date         time.Time
	IsDayOffFunc func(t time.Time) bool
	LinkFunc     func(t time.Time) string
	SelectedDate time.Time
}

var DefaultWeeekLabels = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

func DefaultIsDayOffFunc(t time.Time) bool {
	return t.Weekday() == time.Sunday || t.Weekday() == time.Saturday
}

func NewCalendar() *Calendar {
	return &Calendar{WeekLabels: DefaultWeeekLabels, IsDayOffFunc: DefaultIsDayOffFunc, Date: time.Now()}
}

func (c *Calendar) Html() string {
	wc := len(c.WeekLabels)
	wd := (int(c.Date.Weekday()) - (c.Date.Day() - 1) + wc*30) % wc
	last := c.Date.AddDate(0, 1, -c.Date.Day()).Day()
	s := "<table>\n"
	s += " <tr><th>" + strings.Join(c.WeekLabels, "</th><th>") + "</th></tr>\n"

	for d := 1; d <= last; d++ {
		var attrs string
		if c.IsDayOffFunc(c.Date.AddDate(0, 0, -c.Date.Day()+d)) {
			attrs = " class='dayoff'"
		}
		url := c.link(d)
		if url != "" {
			s += "<td><span" + attrs + "><a href='" + url + "'>" + fmt.Sprint(d) + "</a></span></td>"
		} else {
			s += fmt.Sprintf("<td><span"+attrs+">%2d</span></td>", d)
		}
		wd = (wd + 1) % wc
		if wd == 0 {
			s += "</tr>\n"
			if d != last {
				s += " <tr>"
			}
		}
	}
	s += strings.Repeat("<td></td>", wc-wd) + "<tr>\n"

	s += "</table>\n"
	return s
}

func (c *Calendar) Markdown() string {
	wc := len(c.WeekLabels)
	wd := (int(c.Date.Weekday()) - (c.Date.Day() - 1) + wc*30) % wc
	last := c.Date.AddDate(0, 1, -c.Date.Day()).Day()

	s := "| " + strings.Join(c.WeekLabels, " | ") + " |\n"
	s += "|" + strings.Repeat("----:|", wc) + "\n"
	s += "|" + strings.Repeat("     |", wd)
	for d := 1; d <= last; d++ {
		url := c.link(d)
		if url != "" {
			s += " [" + fmt.Sprint(d) + "](" + url + ") |"
		} else {
			s += fmt.Sprintf("  %2d |", d)
		}
		wd = (wd + 1) % wc
		if wd == 0 {
			s += "\n"
			if d != last {
				s += "|"
			}
		}
	}
	s += strings.Repeat("     |", wc-wd)

	return s
}

func (c *Calendar) String() string {
	return c.Markdown()
}

func (c Calendar) link(d int) string {
	if c.LinkFunc == nil {
		return ""
	}
	return c.LinkFunc(c.Date.AddDate(0, 0, -c.Date.Day()+d))
}
