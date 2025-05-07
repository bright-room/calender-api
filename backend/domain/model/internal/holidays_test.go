package internal_test

import (
	"testing"

	"net.bright-room.dev/calender-api/domain/model/internal"

	"github.com/stretchr/testify/assert"
)

func TestHolidays_Get(t *testing.T) {
	t.Run("指定された日付が祝日に含まれている場合取得できる", func(t *testing.T) {
		holidays := holidaysDummyData

		date1 := parseDate("2025-01-01")
		assert.True(t, holidays.Contains(internal.Reference(date1)))

		date2 := parseDate("2025-01-13")
		assert.True(t, holidays.Contains(internal.Reference(date2)))

		date3 := parseDate("2025-01-14")
		assert.False(t, holidays.Contains(internal.Reference(date3)))
	})
}

func TestHolidays_Contains(t *testing.T) {
	t.Run("指定された日付が祝日に含まれている", func(t *testing.T) {
		holidays := holidaysDummyData

		date1 := parseDate("2025-01-01")
		assert.True(t, holidays.Contains(internal.Reference(date1)))

		date2 := parseDate("2025-01-13")
		assert.True(t, holidays.Contains(internal.Reference(date2)))

		date3 := parseDate("2025-01-14")
		assert.False(t, holidays.Contains(internal.Reference(date3)))
	})
}

func TestHolidays_Concat(t *testing.T) {
	t.Run("祝日リストを結合できる", func(t *testing.T) {
		holidays1 := internal.Holidays{
			internal.NewNationalHoliday(parseDate("2025-01-01"), "aaa"),
		}

		holidays2 := internal.Holidays{
			internal.NewNationalHoliday(parseDate("2025-01-02"), "aaa"),
		}

		concat := holidays1.Concat(holidays2)
		assert.Equal(t, 2, len(concat))
	})
}
