package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/xssnick/tonutils-go/address"
)

func AddrCommand() *cli.Command {
	return &cli.Command{
		Name:  "addr",
		Usage: "Convert raw hex address to bounceable format",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "raw",
				Usage:    "Raw hex address (32 bytes)",
				Required: true,
			},
			&cli.IntFlag{
				Name:  "workchain",
				Value: 0,
				Usage: "Workchain ID (default 0)",
			},
		},
		Action: func(c *cli.Context) error {
			rawHex := c.String("raw")
			workchain := int8(c.Int("workchain"))

			data, err := hex.DecodeString(rawHex)
			if err != nil {
				return fmt.Errorf("❌ Invalid hex string: %w", err)
			}

			if len(data) != 32 {
				return fmt.Errorf("❌ Address must be 32 bytes (got %d)", len(data))
			}

			addr := address.NewAddress(byte(workchain), byte(workchain), data)
			fmt.Println("✅ Bounceable address:", addr.String())
			return nil
		},
	}
}
