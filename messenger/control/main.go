package control

import "github.com/chiraponkub/DPU-SosApp-v.1.git/db"

type ConController struct {
	GORMFactory *db.GORMFactory
	Access      *db.Access
}
