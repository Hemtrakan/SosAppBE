package control

import "github.com/Hemtrakan/SosAppBE.git/db"

type ConController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
