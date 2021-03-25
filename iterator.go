package timeiterator

import "time"

type TimeIterator struct {
	start time.Time
	end   time.Time
}

// New will create a new TimeIterator initialized with the provided values.
func New(start time.Time, end time.Time) *TimeIterator {
	return &TimeIterator{
		start: start,
		end:   end,
	}
}

// Days will iterate over the days between the two time objects
// that were provided when the iterator was created
// (including the dates of start and end).
// If the start comes after the end date, the iteration will go backwards.
// If the date of the start and end objects si the same, the callback will be called
// with the end time object.
func (itr *TimeIterator) Days(callback func(time.Time) bool) {
	if dateEqual(itr.start, itr.end) {
		// If the start and end dates are the same,
		// execute the callback with the end time.
		callback(itr.end)
		return
	}

	var dayDiff int
	var diffFunc func(*time.Time, time.Time) bool
	if itr.start.After(itr.end) {
		dayDiff = -1
		diffFunc = (*time.Time).After
	} else {
		dayDiff = 1
		diffFunc = (*time.Time).Before
	}

	for d := itr.start; diffFunc(&d, itr.end) || dateEqual(d, itr.end); d = d.AddDate(0, 0, dayDiff) {
		doContinue := callback(d)
		if !doContinue {
			return
		}
	}
}

func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// MustParseTime is a utility function wrapping time.Parse;
// will panic on error.
func MustParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

const YearMonthDayFormat = "2006-01-02"
