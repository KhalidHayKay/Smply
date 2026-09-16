package main

import (
	"context"
	"log"
	"smply/config"
	"smply/migrations"

	"github.com/khalidhaykay/cmdforge"
	goosecmd "github.com/khalidhaykay/cmdforge/goose"
)

func main() {
	config.LoadEnv()

	migrator, err := goosecmd.New(config.Env.DbUrl, migrations.FS)
	if err != nil {
		log.Fatal(err)
	}

	cli := cmdforge.New()

	cli.Use(migrator)

	cli.Start(context.Background())
}
