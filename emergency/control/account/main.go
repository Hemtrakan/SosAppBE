package account

import "github.com/chiraponkub/DPU-SosApp-v.1.git/db"

type AccountController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
