package main

import (
	"github.com/hackhack-Geek-vol6/backend/src/frameworks/echo"
	"github.com/hackhack-Geek-vol6/backend/src/server"
)

//	@title						Hack Hack Backend API
//	@version					0.1.0
//	@description			Hack Hack Backend API serice
//	@termsOfService	　https://api.seafood-dev.com

//	@contact.name			murasame29
//	@contact.url			https://twitter.com/fresh_salmon256
//	@contact.email		oogiriminister@gmail.com

//	@license.name			No-license

// @host							api.seafood-dev.com
// @BasePath					/v1
func main() {
	handler := echo.NewEchoServer()
	server.New().Run(handler)
}
