package control

import "accounts/db"

type Controller struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
