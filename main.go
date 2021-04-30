package main

import (
	"login-go/app"
)

func main() {
	app := new(app.App)
	app.Initialize()
	app.SetRouters()
	app.Run(":3000")

}
