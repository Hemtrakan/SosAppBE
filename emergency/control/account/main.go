package account

import (
	"emergency/db"
)

type Controller struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
