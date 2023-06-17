package timezone

import "time"

var Now = time.Now

func FreezeWithTime(timeMs int64) {
	Now = func() time.Time {
		return time.Unix(timeMs/1_000, (timeMs%1_000)*1_000_000)
	}
}

func Unfreeze() {
	Now = time.Now
}
