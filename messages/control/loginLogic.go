package control

import (
	"github.com/chiraponkub/DPU-SosApp-v.1.git/db/structure"
	singin "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singin/request"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/token"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/verify"
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
