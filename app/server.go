package app

import (
	"ginTemplate/app/di"
	"ginTemplate/app/router"
)

func StartGIN() {
	init := di.Init()
	app := router.Init(init)
	app.Run(":8080")
}
