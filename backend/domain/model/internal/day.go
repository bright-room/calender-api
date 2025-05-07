package internal

import "time"

type Day interface {
	IsBusiness() bool
	IsWeekend() bool
	IsHoliday() bool

	Summary() string
	Date() time.Time
	Weekday() time.Weekday
	Description() string
}
