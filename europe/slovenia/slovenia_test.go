package slovenia_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	country "github.com/varrcan/workcalendar/europe/slovenia"
)

var (
	holiday1   = time.Date(2020, time.April, 27, 0, 0, 0, 0, time.UTC)
	holiday2   = time.Date(2020, time.May, 2, 0, 0, 0, 0, time.UTC)
	notHoliday = time.Date(2020, time.May, 17, 0, 0, 0, 0, time.UTC)
)

func TestIsHoliday(t *testing.T) {
	assert.True(t, country.IsHoliday(holiday1))
	assert.True(t, country.IsHoliday(holiday2))
}

func TestGetHoliday(t *testing.T) {
	holiday, err := country.GetHoliday(holiday1)
	v := "Day of Uprising Against Occupation"
	assert.NoError(t, err)
	assert.Equal(t, v, holiday.Name)
}

func TestGetHoliday_NotHoliday(t *testing.T) {
	_, err := country.GetHoliday(notHoliday)
	assert.Error(t, err)
}

func TestIsWorkingDay(t *testing.T) {
	assert.False(t, country.IsWorkingDay(holiday1))
	assert.False(t, country.IsWorkingDay(holiday2))
}
