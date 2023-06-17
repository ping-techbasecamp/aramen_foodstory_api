package errs

import (
	"aramen-api/pkg/utils/timezone"
	"fmt"
	"strconv"
	"time"
)

type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	s := fmt.Sprint(time.Time(t).UnixNano())
	return []byte(s), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	n, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(n/1_000_000_000, n%1_000_000_000))
	return nil
}

func TimestampNow() Timestamp {
	return Timestamp(timezone.Now().UTC())
}
