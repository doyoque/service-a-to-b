package conf

import "time"

var (
	Now = func() time.Time {
		return time.Now().UTC()
	}
)
