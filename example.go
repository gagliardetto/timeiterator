package timeiterator

import (
	"fmt"
	"time"
)

func example() {
	start := MustParseTime(YearMonthDayFormat, "2009-11-10")
	end := MustParseTime(YearMonthDayFormat, "2021-03-25")

	iter := New(start, end)
	iter.Days(func(t time.Time) bool {
		fmt.Println(t.Format(YearMonthDayFormat))
		return true
	})
}
