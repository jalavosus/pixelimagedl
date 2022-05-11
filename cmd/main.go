package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "pixel image downloader",
		Authors: []*cli.Author{
			{Name: "jalavosus", Email: "alavosus.james@gmail.com"},
		},
		Commands: []*cli.Command{
			&downloadCmd,
			&listCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
