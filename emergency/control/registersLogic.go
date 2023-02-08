package control

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/common"
	rdbmsstructure "github.com/chiraponkub/DPU-SosApp-v.1.git/db/structure"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singup/request"
	response "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singup/response"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/verify"
	"math/rand"
	"strconv"
	"time"
)

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func (ctrl ConController) SentOTPLogic(req *request.PhoneNumber) (res response.OTP, Error error) {
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

	newReq := rdbmsstructure.OTP{
		PhoneNumber: req.PhoneNumber,
		Key:         OTP,
		VerifyCode:  VerifyCode,
		Expired:     time.Now().Add(time.Minute * 3).Add(time.Hour * 7),
		Active:      true,
	}

	err = ctrl.Access.RDBMS.SendOTPDB(newReq)
	if err != nil {
		Error = err
		return
	}
	res.OTP = OTP
	res.VerifyCode = VerifyCode
	return
}

func (ctrl ConController) VerifyOTPLogic(req *request.OTP) (Error error) {
	checkNumber, err := common.CheckPhoneNumber(req.PhoneNumber)
	if !checkNumber {
		Error = err
		return
	}
	checkOTPLen, err := common.CheckOTPLen(req.OTP)
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

func (ctrl ConController) CreateUserLogin(req *request.Account) (Error error) {

	db := rdbmsstructure.OTP{
		PhoneNumber: req.PhoneNumber,
		Key:         req.Key,
		VerifyCode:  req.VerifyCode,
	}

	err := ctrl.Access.RDBMS.UpdateOTPDB(db)
	if err != nil {
		Error = err
		return
	}

	if req.Password != req.ConfirmPassword {
		Error = errors.New("รหัสผ่านไม่ตรงกัน")
		return
	}

	hashPass, err := verify.Hash(req.Password)

	newReq := rdbmsstructure.Users{
		PhoneNumber: req.PhoneNumber,
		Password:    string(hashPass),
		Firstname:   req.FirstName,
		Lastname:    req.LastName,
		Email:       &req.Email,
		Birthday:    time.Time{},
		//Birthday:    req.Birthday,
		Gender:       req.Gender,
		ImageProfile: nil,
		DeletedBy:    nil,
		Workplace:    nil,
		AddressID:    1,
		RoleID:       1,
	}

	err = ctrl.Access.RDBMS.CreateUserDB(newReq)
	if err != nil {
		Error = err
		return
	}
	return
}
