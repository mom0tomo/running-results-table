package main

import (
	"github.com/mom0tomo/running-results-table/internal/db"
	"github.com/mom0tomo/running-results-table/internal/webapp"
)

func main() {
	database := db.New()

	webapp.StartServer(&database)
}
