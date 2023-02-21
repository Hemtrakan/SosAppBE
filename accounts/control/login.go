package control

import (
	"accounts/db/structure"
	singin "accounts/restapi/model/singin/request"
	"accounts/utility/token"
	"accounts/utility/verify"
	"gorm.io/gorm"
)

func (ctrl Controller) LoginLogic(request *singin.Login, ip, system string) (Token string, Error error) {

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
	tokenRes, err := token.CreateToken(account.ID, roleId.Name)
	if err != nil {
		Error = err
		return
	}

	Token = tokenRes

	return
}
