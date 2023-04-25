package main

import (
	"github.com/ilhm-rai/mygram/pkg/config"
)

func main() {
	configuration := config.New()
	database := config.NewPostgresDatabase(configuration)

	_ = database
}
