package main

import (
	"fmt"

	"github.com/binzume/go-calendar"
)

func main() {
	fmt.Println(calendar.NewCalendar().Markdown())
}
