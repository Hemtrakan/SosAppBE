package account

import "hotline/db"

type AccountController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
