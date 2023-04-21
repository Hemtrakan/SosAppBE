package common

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CheckPhoneNumber(req string) (res bool, Error error) {
	PhoneNumber, err := regexp.MatchString("^[0-9]{10}$", req)
	if !PhoneNumber {
		Error = errors.New("Username Invalid. : 10 Numbers 0-9")
		return
	}
	if err != nil {
		Error = err
		return
	}

	res = PhoneNumber
	return
}

func CheckOTPLen(OTP int) (res bool, Error error) {
	res = false
	if len(strconv.Itoa(OTP)) == 4 {
		res = true
	}
	Error = errors.New("OTP Invalid")
	return
}

func changeTimeLayout(layout string) string {
	layout = strings.ReplaceAll(layout, "yyyy", "2006")
	layout = strings.ReplaceAll(layout, "mm", "01")
	layout = strings.ReplaceAll(layout, "dd", "02")
	layout = strings.ReplaceAll(layout, "hh", "15")
	layout = strings.ReplaceAll(layout, "mi", "04")
	layout = strings.ReplaceAll(layout, "ss", "05")
	layout = strings.ReplaceAll(layout, "SSS", "000")

	return layout
}

/*
input layout:

	yyyy-mm-dd hh:mi:ss
	yyyy-mm-dd hh:mi:ss.SSS
	yyyy-mm-ddThh:mi:ss.SSS
	yyyy-mm-ddThh:mi:ss.SSSZ0700
	yyyy-mm-ddThh:mi:ssZ07:00
*/
func StringToTime(layout string, dateTime string) (time.Time, error) {
	layout = changeTimeLayout(layout)
	return time.Parse(layout, dateTime)
}

/*
input layout:

	yyyy-mm-dd hh:mi:ss
	yyyy-mm-dd hh:mi:ss.SSS
	yyyy-mm-ddThh:mi:ss.SSS
	yyyy-mm-ddThh:mi:ss.SSSZ0700
	yyyy-mm-ddThh:mi:ssZ07:00
*/
func TimeToString(layout string, dateTime time.Time) string {
	var dt string

	if !dateTime.IsZero() {
		layout = changeTimeLayout(layout)
		dt = dateTime.Format(layout)
	}

	return dt
}

func IntToString(data int) string {
	return strconv.FormatInt(int64(data), 10)
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func GetStringFloat(value string) string {
	if s, err := strconv.ParseFloat(value, 64); err == nil {
		return fmt.Sprintf("%.2f", s)
	} else {
		return value
	}
}

func SetStartDay(dateTime time.Time) time.Time {
	var startDT time.Time

	if !dateTime.IsZero() {
		startDT = time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(),
			0, 0, 0, 0, dateTime.Location())
	}

	return startDT
}

func SetEndDay(dateTime time.Time) time.Time {
	var endDT time.Time

	if !dateTime.IsZero() {
		endDT = time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(),
			23, 59, 59, 999999999, dateTime.Location())
	}

	return endDT
}

func GetString(value interface{}) string {
	v, stringNotNil := value.(string)
	if stringNotNil {
		return v
	} else {
		return ""
	}
}
