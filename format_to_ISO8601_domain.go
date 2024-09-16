package sumaregi

import (
	"time"
)

func FormatToISO8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-07:00")
}
