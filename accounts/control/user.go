package control

import (
	rdbmsstructure "accounts/db/structure"
	"accounts/restapi/model/singup/request"
	resUser "accounts/restapi/model/user/response"
	//reqUser "accounts/restapi/model/user/request"
	"accounts/utility/verify"
	"errors"
	"gorm.io/gorm"
)

func (ctrl Controller) GetUser(id uint) (res *resUser.UserProfile, Error error) {
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
	res = &resUser.UserProfile{
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

func (ctrl Controller) PostUser(req *request.Account) (Error error) {
	//otp := rdbmsstructure.OTP{
	//	PhoneNumber: req.PhoneNumber,
	//	Key:         req.Verify.OTP,
	//	VerifyCode:  req.Verify.VerifyCode,
	//}
	//
	//err := ctrl.Access.RDBMS.UpdateOTPDB(otp)
	//if err != nil {
	//	Error = err
	//	return
	//}

	if req.Password != req.ConfirmPassword {
		Error = errors.New("รหัสผ่านไม่ตรงกัน")
		return
	}

	hashPass, err := verify.Hash(req.Password)

	roleModel := rdbmsstructure.Role{
		Name: "user",
	}
	role, err := ctrl.Access.RDBMS.GetRoleDBByName(roleModel)
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

	err = ctrl.Access.RDBMS.CreateUserDB(newReq)
	if err != nil {
		Error = err
		return
	}
	return
}
