package internal_test

import (
	"time"

	"net.bright-room.dev/calender-api/domain/model/internal"
)

func parseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		panic(err)
	}
	return t
}

var holidaysDummyData = internal.Holidays{
	internal.NewNationalHoliday(parseDate("2025-01-01"), "元日"),
	internal.NewNationalHoliday(parseDate("2025-01-13"), "成人の日"),
	internal.NewNationalHoliday(parseDate("2025-02-11"), "建国記念の日"),
	internal.NewNationalHoliday(parseDate("2025-02-23"), "天皇誕生日"),
	internal.NewNationalHoliday(parseDate("2025-02-24"), "休日"),
	internal.NewNationalHoliday(parseDate("2025-03-20"), "春分の日"),
	internal.NewNationalHoliday(parseDate("2025-04-29"), "昭和の日"),
	internal.NewNationalHoliday(parseDate("2025-05-03"), "憲法記念日"),
	internal.NewNationalHoliday(parseDate("2025-05-04"), "みどりの日"),
	internal.NewNationalHoliday(parseDate("2025-05-05"), "こどもの日"),
	internal.NewNationalHoliday(parseDate("2025-05-06"), "休日"),
	internal.NewNationalHoliday(parseDate("2025-07-21"), "海の日"),
	internal.NewNationalHoliday(parseDate("2025-08-11"), "山の日"),
	internal.NewNationalHoliday(parseDate("2025-09-15"), "敬老の日"),
	internal.NewNationalHoliday(parseDate("2025-09-23"), "秋分の日"),
	internal.NewNationalHoliday(parseDate("2025-10-13"), "スポーツの日"),
	internal.NewNationalHoliday(parseDate("2025-11-03"), "文化の日"),
	internal.NewNationalHoliday(parseDate("2025-11-23"), "勤労感謝の日"),
	internal.NewNationalHoliday(parseDate("2025-11-24"), "休日"),
}
