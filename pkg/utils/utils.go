package utils

import (
	"aramen-api/pkg/utils/timezone"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	DateKey = "date"
)

func BangkokNow() time.Time {
	return timezone.Now().In(timezone.BangkokLocation())
}

func IsEqualID(id uint, idRequest any) bool {
	strIDRequest, ok := idRequest.(string)
	if !ok {
		return false
	}

	uintidRequest, err := strconv.ParseUint(strIDRequest, 10, 64)
	if err != nil {
		return false
	}

	return uint(uintidRequest) == id
}

func ParseDate(c *gin.Context) (time.Time, error) {
	date := c.Query(DateKey)
	if date == "" {
		return time.Time{}, errors.New("date is required")
	}

	result, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}

	return result, nil
}
