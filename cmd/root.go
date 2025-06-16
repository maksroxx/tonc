package cmd

import (
	"log"
	"os"

	"github.com/maksroxx/tonc/build"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name:    "tonc",
		Usage:   "FunC compiler tool",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Compile FunC contracts (single or all in folder)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "contract",
						Aliases: []string{"c"},
						Usage:   "Path to single contract file to compile",
					},
					&cli.StringFlag{
						Name:    "src",
						Aliases: []string{"s"},
						Usage:   "Source directory with .fc files (used if --contract not set)",
						Value:   "./contracts",
						EnvVars: []string{"TONC_SRC"},
					},
					&cli.StringFlag{
						Name:    "out",
						Aliases: []string{"o"},
						Usage:   "Output directory for build artifacts",
						Value:   "build",
					},
					&cli.BoolFlag{
						Name:    "boc",
						Usage:   "Save .cell.boc files",
						Value:   true,
						Aliases: []string{"B"},
					},
					&cli.BoolFlag{
						Name:    "json",
						Usage:   "Save compiled JSON files",
						Value:   true,
						Aliases: []string{"J"},
					},
					&cli.BoolFlag{
						Name:    "hex",
						Usage:   "Include hex string in JSON output",
						Value:   true,
						Aliases: []string{"H"},
					},
					&cli.BoolFlag{
						Name:  "verbose",
						Usage: "Enable verbose logging",
						Value: false,
					},
				},
				Action: build.BuildAction,
			},
			AddrCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("❌ %v", err)
	}
}
