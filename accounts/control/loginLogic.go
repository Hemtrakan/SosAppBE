package control

import (
	"github.com/Hemtrakan/SosAppBE.git/db/structure"
	singin "github.com/Hemtrakan/SosAppBE.git/restapi/model/singin/request"
	"github.com/Hemtrakan/SosAppBE.git/utility/token"
	"github.com/Hemtrakan/SosAppBE.git/utility/verify"
	"gorm.io/gorm"
)

func (ctrl ConController) LoginLogic(request *singin.Login) (Token string, Error error) {

	db := structure.Users{
		PhoneNumber: request.Username,
	}

	account, err := ctrl.Access.RDBMS.GetAccountDB(db)
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
			//ID: account.RoleID,
		},
	}

	roleId, err := ctrl.Access.RDBMS.GetRoleDBById(roleStr)
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
