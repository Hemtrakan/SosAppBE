package common

import (
	"errors"
	"regexp"
	"strconv"
)

func CheckPhoneNumber(req string) (res bool, Error error) {
	PhoneNumber, err := regexp.MatchString("^[0-9]{10}$", req)
	if !PhoneNumber {
		Error = errors.New("PhoneNumber Invalid. : 10 Numbers 0-9")
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
