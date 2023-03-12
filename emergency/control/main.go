package control

import (
	"emergency/db"
	"emergency/httpclient"
)

type Controller struct {
	Access     *db.Access
	HttpClient httpclient.HttpClient
}
