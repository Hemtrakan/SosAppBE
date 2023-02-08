package control

import "hotline/db"

type ConController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
