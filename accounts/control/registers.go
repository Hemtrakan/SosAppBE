package control

import (
	"accounts/common"
	rdbmsstructure "accounts/db/structure"
	"accounts/restapi/model/singup/request"
	response "accounts/restapi/model/singup/response"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func (ctrl Controller) SentOTPLogic(req *request.PhoneNumber) (res response.OTP, Error error) {
	Check, err := common.CheckPhoneNumber(req.PhoneNumber)
	if !Check {
		Error = errors.New("PhoneNumber Invalid. : 10 Numbers 0-9")
		return
	}
	if err != nil {
		Error = err
		return
	}

	var OTP int
	for {
		OTP = rangeIn(0000, 9999)
		if len(strconv.Itoa(OTP)) == 4 {
			break
		}
	}

	rand.Seed(time.Now().UnixNano())
	VerifyCode := ""
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	i := 0
	for {
		c := charset[rand.Intn(len(charset))]
		VerifyCode = VerifyCode + string(c)
		if i == 4 {
			break
		}
		i++
	}
	strOTP := strconv.Itoa(OTP)
	newReq := rdbmsstructure.OTP{
		PhoneNumber: req.PhoneNumber,
		Key:         strOTP,
		VerifyCode:  VerifyCode,
		Expired:     time.Now().Add(time.Minute * 3).Add(time.Hour * 7),
		Active:      true,
	}

	err = ctrl.Access.RDBMS.SendOTPDB(newReq)
	if err != nil {
		Error = err
		return
	}
	res.OTP = strOTP
	res.VerifyCode = VerifyCode
	return
}

func (ctrl Controller) VerifyOTPLogic(req *request.OTP) (Error error) {
	checkNumber, err := common.CheckPhoneNumber(req.PhoneNumber)
	if !checkNumber {
		Error = err
		return
	}

	OTP, err := strconv.Atoi(req.OTP)
	if err != nil {
		Error = err
		return
	}

	checkOTPLen, err := common.CheckOTPLen(OTP)
	if !checkOTPLen {
		Error = err
		return
	}
	newReq := rdbmsstructure.OTP{
		PhoneNumber: req.PhoneNumber,
		Key:         req.OTP,
		VerifyCode:  req.VerifyCode,
	}

	res, err := ctrl.Access.RDBMS.GetOTPDb(newReq)
	if err != nil {
		Error = err
		return
	}
	t1 := time.Now().Add(time.Hour * 7)
	t2 := res.Expired
	diff := t2.Sub(t1).Seconds()
	if int(diff) < 0 {
		Error = errors.New("รหัส OTP หมดอายุ")
		return
	}

	if res.Active != true || res.VerifyCode != req.VerifyCode || res.PhoneNumber != req.PhoneNumber || res.Key != req.OTP {
		Error = errors.New("รหัส OTP ไม่ถูกต้อง")
		return
	}

	return
}
