package control

import (
	"accounts/db/structure"
)

func (ctrl Controller) GetLogLogin() (res []structure.LogLogin, Error error) {
	data, err := ctrl.Access.RDBMS.GetLogLogin()
	if err != nil {
		Error = err
		return
	}

	res = data
	return
}
