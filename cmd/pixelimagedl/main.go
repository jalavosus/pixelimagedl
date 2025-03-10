package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var (
	author = map[string]string{"Name": "jalavosus", "Email": "alavosus.james@gmail.com"}
)

func main() {
	app := &cli.Command{
		Name:    "pixelimagedl",
		Authors: []any{author},
		Commands: []*cli.Command{
			&downloadCmd,
			&listCmd,
		},
	}

	ctx := context.TODO()

	if err := app.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
