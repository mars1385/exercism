package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 2
	Third  WeekSchedule = 3
	Fourth WeekSchedule = 4
	Last   WeekSchedule = 5
	Teenth WeekSchedule = 6
)

func isValidTeenth(day int) bool {
	return day >= 13 && day <= 19
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {

	time := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	endDate := time.AddDate(0, 1, 0)

	day := 0
	counter := 0
	for date := time; !date.After(endDate) && !date.Equal(endDate); date = date.AddDate(0, 0, 1) {

		if date.Weekday() == wDay {
			day = date.Day()
			counter++
		}

		if wSched <= Fourth && date.Weekday() == wDay && counter == int(wSched) {
			return date.Day()
		}

		if wSched == Teenth && date.Weekday() == wDay && isValidTeenth(date.Day()) {
			return date.Day()
		}

	}

	return day
}
