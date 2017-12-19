package main

import "github.com/Attsun1031/jobnetes/cmd/dbmigration/app"

func main() {
	app.NewDbmigrationCommand().Execute()
}
