package control

import (
	rdbmsstructure "accounts/db/structure"
	"accounts/restapi/model/singup/request"
	user "accounts/restapi/model/user/response"
	"accounts/utility/verify"
	"errors"
	"gorm.io/gorm"
	"time"
)

func (ctrl Controller) GetUser(id uint) (res *user.UserProfile, Error error) {
	req := rdbmsstructure.Users{
		Model: gorm.Model{
			ID: id,
		},
	}

	data, err := ctrl.Access.RDBMS.GetUserByID(req)
	if err != nil {
		Error = err
		return
	}

	res = &user.UserProfile{
		PhoneNumber: data.PhoneNumber,
		FirstName:   data.Firstname,
		LastName:    data.Lastname,
		Birthday:    data.Birthday,
		Gender:      data.Gender,
		IDCard:      data.IDCard.IDCardText,
		RoleID:      data.Role.Name,
	}
	return
}

func (ctrl Controller) PostUser(req *request.Account) (Error error) {
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
