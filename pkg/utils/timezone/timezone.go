package timezone

import "time"

var loc *time.Location

func init() {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	loc = location
}

func BangkokLocation() *time.Location {
	return loc
}
