package internal

import (
	"log"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var ClientStore *store.ClientStore
var Oauth2Server *server.Server

func InitOauth2() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	ClientStore = store.NewClientStore()
	manager.MapClientStorage(ClientStore)

	Oauth2Server = server.NewDefaultServer(manager)
	Oauth2Server.SetAllowGetAccessRequest(true)
	Oauth2Server.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	Oauth2Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	Oauth2Server.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
}
