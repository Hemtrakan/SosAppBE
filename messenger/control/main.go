package control

import (
	"messenger/db"
	"messenger/httpclient"
)

type Controller struct {
	Access     *db.Access
	HttpClient httpclient.HttpClient
}
