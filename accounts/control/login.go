package control

import (
	"accounts/constant"
	"accounts/db/structure"
	singin "accounts/restapi/model/singin/request"
	"accounts/utility/pointer"
	"accounts/utility/token"
	"accounts/utility/verify"
	"errors"
	"gorm.io/gorm"
)

func (ctrl Controller) Login(request *singin.Login, ip, system string) (Token, description string, Error error) {

	db := structure.Users{
		PhoneNumber: request.Username,
	}

	account, err := ctrl.Access.RDBMS.GetUserByPhone(db)
	if err != nil {
		Error = err
		return
	}

	checkPass := verify.VerifyPassword(account.Password, request.Password)
	if checkPass != nil {
		Error = err
		return
	}

	if pointer.GetStringValue(account.IDCard.Description) != "" && !account.IDCard.Verify {
		description = pointer.GetStringValue(account.IDCard.Description)
		return
	}

	if account.IDCard.Verify == false {
		Error = errors.New("บัญชีของท่านยังไม่ถูกยืนยันตัวตน กรุณาเข้าสู่ระบบอีกครั้งหลังจากยืนยันตัวตนสำเร็จ")
		return
	}

	roleStr := structure.Role{
		Model: gorm.Model{
			ID: account.RoleID,
		},
	}

	roleId, err := ctrl.Access.RDBMS.GetRoleById(roleStr)
	if err != nil {
		Error = err
		return
	}

	log := structure.LogLogin{
		UserID:      account.ID,
		System:      system,
		IP:          ip,
		Status:      constant.SuccessMsg,
		Description: "Role Login : " + roleId.Name,
	}

	err = ctrl.Access.RDBMS.LogLogin(log)
	if err != nil {
		Error = err
		return
	}

	tokenRes, err := token.CreateToken(account.ID, roleId.Name)
	if err != nil {
		Error = err
		return
	}

	Token = tokenRes

	return
}
