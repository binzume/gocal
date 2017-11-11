package calendar

import (
	"strings"
	"testing"
	"time"
)

func TestCalendar(t *testing.T) {
	t.Log(NewCalendar().Html())
	t.Log(NewCalendar().Markdown())
	t.Log(NewCalendar().String())
}

func TestMarkdown(t *testing.T) {
	cal := NewCalendar()
	cal.Date, _ = time.Parse("2006-01-02", "2017-12-01")

	{
		expected := strings.Replace(`| Sun | Mon | Tue | Wed | Thu | Fri | Sat |
		|----:|----:|----:|----:|----:|----:|----:|
		|     |     |     |     |     |   1 |   2 |
		|   3 |   4 |   5 |   6 |   7 |   8 |   9 |
		|  10 |  11 |  12 |  13 |  14 |  15 |  16 |
		|  17 |  18 |  19 |  20 |  21 |  22 |  23 |
		|  24 |  25 |  26 |  27 |  28 |  29 |  30 |
		|  31 |     |     |     |     |     |     |`, "\t", "", -1)
		actual := cal.Markdown()
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}

	}
	{
		cal.WeekLabels = []string{"sun", "sun", "Sat", "Sat", "Sat", "Sat", "Sun"}
		expected := strings.Replace(`| sun | sun | Sat | Sat | Sat | Sat | Sun |
		|----:|----:|----:|----:|----:|----:|----:|
		|     |     |     |     |     |   1 |   2 |
		|   3 |   4 |   5 |   6 |   7 |   8 |   9 |
		|  10 |  11 |  12 |  13 |  14 |  15 |  16 |
		|  17 |  18 |  19 |  20 |  21 |  22 |  23 |
		|  24 |  25 |  26 |  27 |  28 |  29 |  30 |
		|  31 |     |     |     |     |     |     |`, "\t", "", -1)
		actual := cal.Markdown()
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}

	}
}

func TestLink(t *testing.T) {
	cal := NewCalendar()
	cal.Date, _ = time.Parse("2006-01-02", "2017-12-01")
	cal.LinkFunc = func(t time.Time) string {
		if t.Day() == 24 {
			return "http://example.com/"
		}
		return ""
	}
	expected := strings.Replace(`| Sun | Mon | Tue | Wed | Thu | Fri | Sat |
		|----:|----:|----:|----:|----:|----:|----:|
		|     |     |     |     |     |   1 |   2 |
		|   3 |   4 |   5 |   6 |   7 |   8 |   9 |
		|  10 |  11 |  12 |  13 |  14 |  15 |  16 |
		|  17 |  18 |  19 |  20 |  21 |  22 |  23 |
		| [24](http://example.com/) |  25 |  26 |  27 |  28 |  29 |  30 |
		|  31 |     |     |     |     |     |     |`, "\t", "", -1)
	actual := cal.Markdown()
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
