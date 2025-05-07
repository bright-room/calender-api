package internal

import "time"

type Weekend struct {
	date time.Time
}

func NewWeekend(date time.Time) Weekend {
	return Weekend{date: date}
}

func (r Weekend) IsBusiness() bool {
	return false
}

func (r Weekend) IsWeekend() bool {
	return true
}

func (r Weekend) IsHoliday() bool {
	return false
}

func (r Weekend) Summary() string {
	return "週末"
}

func (r Weekend) Date() time.Time {
	return r.date
}

func (r Weekend) Weekday() time.Weekday {
	return r.date.Weekday()
}

func (r Weekend) Description() string {
	if r.date.Weekday() == time.Saturday {
		return "土曜日"
	}

	return "日曜日"
}
