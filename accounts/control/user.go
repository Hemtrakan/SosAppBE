package control

import (
	_image "accounts/assets/image"
	rdbmsstructure "accounts/db/structure"
	reqSingUp "accounts/restapi/model/singup/request"
	"accounts/restapi/model/user/request"
	resUser "accounts/restapi/model/user/response"
	"accounts/utility/pointer"
	"accounts/utility/verify"
	"strconv"
	"strings"
	"time"

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
		Email:        data.Email,
		Birthday:     data.Birthday,
		Gender:       data.Gender,
		ImageProfile: *data.ImageProfile,
		Workplace:    pointer.GetStringValue(data.Workplace),
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

func (ctrl Controller) GetUserList() (res []resUser.UserRes, Error error) {

	data, err := ctrl.Access.RDBMS.GetUserList()
	if err != nil {
		Error = err
		return
	}

	for _, user := range data {
		objectUser := resUser.UserRes{
			ID:           strconv.Itoa(int(user.ID)),
			PhoneNumber:  user.PhoneNumber,
			FirstName:    user.Firstname,
			LastName:     user.Lastname,
			Email:        user.Email,
			Birthday:     user.Birthday,
			Gender:       user.Gender,
			ImageProfile: pointer.GetStringValue(user.ImageProfile),
			Workplace:    pointer.GetStringValue(user.Workplace),
			IdCard: resUser.IdCard{
				TextIDCard: user.IDCard.TextIDCard,
				PathImage:  user.IDCard.PathImage,
				Verify:     false,
			},
			Address: resUser.Address{
				Address:     user.Address.Address,
				SubDistrict: user.Address.SubDistrict,
				District:    user.Address.District,
				Province:    user.Address.Province,
				PostalCode:  user.Address.PostalCode,
				Country:     user.Address.Country,
			},
		}

		res = append(res, objectUser)
	}

	return
}

func (ctrl Controller) PostUser(req *reqSingUp.SingUp) (resUsers rdbmsstructure.Users, Error error) {
	checkUserData := rdbmsstructure.Users{
		PhoneNumber: req.PhoneNumber,
	}
	checkUser, err := ctrl.Access.RDBMS.GetUserByPhone(checkUserData)
	if err != nil {
		Error = err
		return
	}
	if checkUser.PhoneNumber == req.PhoneNumber {
		Error = errors.New("This phone number has already been register.")
		return
	}

	if req.Password != req.ConfirmPassword {
		Error = errors.New("รหัสผ่านไม่ตรงกัน")
		return
	}

	hashPass, err := verify.Hash(req.Password)

	roleModel := rdbmsstructure.Role{
		Model: gorm.Model{
			ID: req.RoleId,
		},
	}
	role, err := ctrl.Access.RDBMS.GetRoleById(roleModel)
	if err != nil {
		Error = err
		return
	}

	date := time.Now()
	if req.Birthday != "" {
		Birthday := strings.Split(req.Birthday, " ")
		date, err = time.Parse("2006-01-02", Birthday[0])
		if err != nil {
			Error = err
			return
		}
	} else {
		Error = errors.New("กรุณาเพิ่มวันเดือนปีเกิด")
		return
	}

	image := ""
	if req.ImageProfile == "" {
		image, err = _image.ImageToBase64()
		if err != nil {
			Error = err
			return
		}
	} else {
		image = req.ImageProfile
	}

	otp := rdbmsstructure.OTP{
		PhoneNumber: req.PhoneNumber,
		Key:         req.Verify.OTP,
		VerifyCode:  req.Verify.VerifyCode,
	}

	err = ctrl.Access.RDBMS.UpdateOTPDB(otp)
	if err != nil {
		Error = err
		return
	}

	newReq := rdbmsstructure.Users{
		PhoneNumber:  req.PhoneNumber,
		Password:     string(hashPass),
		Firstname:    req.FirstName,
		Lastname:     req.LastName,
		Email:        req.Email,
		Birthday:     date,
		Gender:       req.Gender,
		ImageProfile: &image,
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

	resUsers.PhoneNumber = req.PhoneNumber
	resUsers.Password = req.Password
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
	var date = time.Time{}
	image := ""
	if req.Gender == "" {
		image = req.ImageProfile
	} else {
		Birthday := strings.Split(req.Birthday, " ")
		date, err = time.Parse("2006-01-02", Birthday[0])
		if err != nil {
			Error = append(Error, err)
			return
		}
		if req.ImageProfile == "" {
			image, err = _image.ImageToBase64()
			if err != nil {
				Error = append(Error, err)
				return
			}
		} else {
			image = req.ImageProfile
		}
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
			Email:        req.Email,
			Birthday:     date,
			Gender:       req.Gender,
			ImageProfile: &image,
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

func (ctrl Controller) ChangePassword(req *request.ChangePassword, userID uint) (Error error) {
	if req.NewPassword != req.ConfirmPassword {
		Error = errors.New("รหัสผ่านไม่ตรงกัน")
		return
	}

	mapData := rdbmsstructure.Users{
		Model: gorm.Model{
			ID: userID,
		},
	}

	userData, err := ctrl.Access.RDBMS.GetUserByID(mapData)
	if err != nil {
		Error = err
		return
	}

	checkPass := verify.VerifyPassword(userData.Password, req.OldPassword)
	if checkPass != nil {
		Error = errors.New("รหัสผ่านไม่ถูกต้อง")
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

	err = ctrl.Access.RDBMS.ChangePassword(Users)
	if err != nil {
		Error = err
	}

	return
}
