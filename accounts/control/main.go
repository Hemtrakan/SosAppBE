package control

import "accounts/db"

type ConController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
