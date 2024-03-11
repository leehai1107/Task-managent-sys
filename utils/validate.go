package utils

import (
	"errors"
	"time"
)

func ValidateDate(startDate DateTime, endDate DateTime) error {
	startTime := time.Time(startDate)
	endTime := time.Time(endDate)

	if endTime.Before(startTime) {
		return errors.New("end date should be greater than start date")
	}

	return nil
}
