package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeRange_DatesUntil(t *testing.T) {
	t.Run("開始日が終了日よりも後の場合エラーになる", func(t *testing.T) {
		begin := time.Date(2025, time.January, 2, 0, 0, 0, 0, JST)
		end := time.Date(2025, time.January, 1, 0, 0, 0, 0, JST)

		tr := TimeRange{Begin: begin, End: end}
		_, err := tr.DatesUntil()

		assert.Error(t, err)
	})

	t.Run("開始日から終了日までの日付リストを作成できる", func(t *testing.T) {
		begin := time.Date(2025, time.January, 1, 0, 0, 0, 0, JST)
		end := time.Date(2025, time.January, 31, 0, 0, 0, 0, JST)

		tr := TimeRange{Begin: begin, End: end}
		actual, _ := tr.DatesUntil()

		assert.Equal(t, 31, len(actual))
	})
}
