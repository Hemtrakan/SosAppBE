package control

import (
	"emergency/db"
)

type ConController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
