package internal_test

import (
	"testing"
	"time"

	"net.bright-room.dev/calender-api/domain/model/internal"

	"github.com/stretchr/testify/assert"
	"net.bright-room.dev/calender-api/extentions/timex"
)

func TestWeekend_Description(t *testing.T) {
	t.Run("2025-04-13は土曜日であると説明文が生成される", func(t *testing.T) {
		date := time.Date(2025, time.April, 12, 0, 0, 0, 0, timex.JST)
		weekend := internal.NewWeekend(date)

		assert.Equal(t, "土曜日", weekend.Description())
	})

	t.Run("2025-04-14は日曜日であると説明文が生成される", func(t *testing.T) {
		date := time.Date(2025, time.April, 13, 0, 0, 0, 0, timex.JST)
		weekend := internal.NewWeekend(date)

		assert.Equal(t, "日曜日", weekend.Description())
	})
}
