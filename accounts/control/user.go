package control

import (
	rdbmsstructure "accounts/db/structure"
	reqSingUp "accounts/restapi/model/singup/request"
	"accounts/restapi/model/user/request"
	resUser "accounts/restapi/model/user/response"
	"strconv"

	//reqUser "accounts/restapi/model/user/request"
	"accounts/utility/verify"
	"errors"
	"gorm.io/gorm"
)

func (ctrl Controller) GetUser(id uint) (res *resUser.UserRes, Error error) {
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
	res = &resUser.UserRes{
		ID:           strconv.Itoa(int(id)),
		PhoneNumber:  data.PhoneNumber,
		FirstName:    data.Firstname,
		LastName:     data.Lastname,
		Email:        *data.Email,
		Birthday:     data.Birthday,
		Gender:       data.Gender,
		ImageProfile: *data.ImageProfile,
		IdCard: resUser.IdCard{
			TextIDCard: data.IDCard.TextIDCard,
			PathImage:  data.IDCard.PathImage,
			Verify:     false,
		},
		Address: resUser.Address{
			Address:     data.Address.Address,
			SubDistrict: data.Address.SubDistrict,
			District:    data.Address.District,
			Province:    data.Address.Province,
			PostalCode:  data.Address.PostalCode,
			Country:     data.Address.Country,
		},
	}
	return
}

func (ctrl Controller) PostUser(req *reqSingUp.SingUp) (Error error) {
	otp := rdbmsstructure.OTP{
		PhoneNumber: req.PhoneNumber,
		Key:         req.Verify.OTP,
		VerifyCode:  req.Verify.VerifyCode,
	}

	err := ctrl.Access.RDBMS.UpdateOTPDB(otp)
	if err != nil {
		Error = err
		return
	}

	if req.Password != req.ConfirmPassword {
		Error = errors.New("รหัสผ่านไม่ตรงกัน")
		return
	}

	hashPass, err := verify.Hash(req.Password)

	roleModel := rdbmsstructure.Role{
		Name: "user",
	}
	role, err := ctrl.Access.RDBMS.GetRoleByName(roleModel)
	if err != nil {
		Error = err
		return
	}

	newReq := rdbmsstructure.Users{
		PhoneNumber:  req.PhoneNumber,
		Password:     string(hashPass),
		Firstname:    req.FirstName,
		Lastname:     req.LastName,
		Email:        &req.Email,
		Birthday:     req.Birthday,
		Gender:       req.Gender,
		ImageProfile: &req.ImageProfile,
		DeletedBy:    nil,
		Workplace:    nil,
		IDCard: rdbmsstructure.IDCard{
			TextIDCard: req.IDCard.TextIDCard,
			PathImage:  req.IDCard.PathImage,
			DeletedBy:  nil,
		},
		Address: rdbmsstructure.Address{
			Address:     req.Address.Address,
			SubDistrict: req.Address.SubDistrict,
			District:    req.Address.District,
			Province:    req.Address.Province,
			PostalCode:  req.Address.PostalCode,
			Country:     req.Address.Country,
			DeletedBy:   nil,
		},
		RoleID: role.ID,
	}

	err = ctrl.Access.RDBMS.PostUser(newReq)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl Controller) PutUser(req *request.UserReq, userID uint) (Error []error) {
	reqUserId := rdbmsstructure.Users{
		Model: gorm.Model{
			ID: userID,
		},
	}
	data, err := ctrl.Access.RDBMS.GetUserByID(reqUserId)
	if err != nil {
		Error = append(Error, err)
		return
	}

	var Users = new(rdbmsstructure.Users)
	var Address = new(rdbmsstructure.Address)
	var IDCard = new(rdbmsstructure.IDCard)
	if req != nil {
		Users = &rdbmsstructure.Users{
			Model: gorm.Model{
				ID: userID,
			},
			Firstname:    req.FirstName,
			Lastname:     req.LastName,
			Email:        &req.Email,
			Birthday:     req.Birthday,
			Gender:       req.Gender,
			ImageProfile: &req.ImageProfile,
			UpdateBy:     &userID,
		}
	}
	if req.Address != nil {
		Address = &rdbmsstructure.Address{
			Model: gorm.Model{
				ID: data.AddressID,
			},
			Address:     req.Address.Address,
			SubDistrict: req.Address.SubDistrict,
			District:    req.Address.District,
			Province:    req.Address.Province,
			PostalCode:  req.Address.PostalCode,
			Country:     req.Address.Country,
			DeletedBy:   nil,
			UpdateBy:    &userID,
		}
	}
	if req.IdCard != nil {
		IDCard = &rdbmsstructure.IDCard{
			Model: gorm.Model{
				ID: data.IDCardID,
			},
			TextIDCard: req.IdCard.TextIDCard,
			PathImage:  req.IdCard.PathImage,
			UpdateBy:   &userID,
		}
	}
	errArr := ctrl.Access.RDBMS.PutUser(Users, Address, IDCard)
	if err != nil {
		Error = errArr
		return
	}
	return
}

func (ctrl Controller) DeleteUser(UserID uint) (Error error) {
	newReq := rdbmsstructure.Users{
		Model: gorm.Model{
			ID: UserID,
		},
	}
	err := ctrl.Access.RDBMS.DeleteUser(newReq)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl Controller) ChangePassword(req *request.UserReq, userID uint) (Error []error) {
	if req.NewPassword != req.ConfirmPassword {
		Error = append(Error, errors.New("รหัสผ่านไม่ตรงกัน"))
		return
	}

	mapData := rdbmsstructure.Users{
		Model: gorm.Model{
			ID: userID,
		},
	}

	userData, err := ctrl.Access.RDBMS.GetUserByPhone(mapData)
	if err != nil {
		Error = append(Error, err)
		return
	}

	checkPass := verify.VerifyPassword(userData.Password, req.OldPassword)
	if checkPass != nil {
		Error = append(Error, errors.New("รหัสผ่านไม่ถูกต้อง"))
		return
	}

	hashPass, err := verify.Hash(req.NewPassword)

	var Users = new(rdbmsstructure.Users)
	if req != nil {
		Users = &rdbmsstructure.Users{
			Model: gorm.Model{
				ID: userID,
			},
			Password: string(hashPass),
		}
	}

	errArr := ctrl.Access.RDBMS.PutUser(Users, nil, nil)
	if err != nil {
		Error = errArr
		return
	}
	return
}
