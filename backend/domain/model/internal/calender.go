package internal

import (
	"time"

	"golang.org/x/xerrors"
	"net.bright-room.dev/calender-api/extentions/timex"
)

type Calendar struct {
	days Days
}

func NewCalendar(holidays Holidays) (Calendar, error) {
	tr := timex.TimeRange{
		Begin: timex.StartOfUnixEpoch(),
		End:   timex.NowDate().AddDate(5, 0, 0),
	}
	until, err := tr.DatesUntil()
	if err != nil {
		return Calendar{}, xerrors.Errorf("failed to create calender: %w", err)
	}

	var calender []Day
	for _, day := range until {
		refDay := Reference(day)

		if holidays.Contains(refDay) {
			holiday, err := holidays.Get(refDay)
			if err != nil {
				return Calendar{}, xerrors.Errorf("failed to create calender: %w", err)
			}

			calender = append(calender, holiday)
			continue
		}

		dayOfWeek := refDay.Weekday()
		if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
			weekend := Weekend{refDay.Date()}
			calender = append(calender, weekend)
			continue
		}

		business := Businessday{refDay.Date()}
		calender = append(calender, business)
	}

	return Calendar{days: calender}, nil
}
