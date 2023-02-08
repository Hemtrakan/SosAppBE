package account

import "github.com/Hemtrakan/SosAppBE.git/db"

type AccountController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
