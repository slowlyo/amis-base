package main

import (
	"amis-base/cmd/amis-base/cmd"
	"amis-base/internal/app"
)

func main() {
	app.Bootstrap()

	cmd.Execute()
}
