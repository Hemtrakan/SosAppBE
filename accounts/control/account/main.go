package account

import "accounts/db"

type AccountController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
