package main

import (
	"encoding/hex"
	"fmt"
	"github.com/civet148/btckeys"
	"github.com/civet148/log"
	cli "github.com/urfave/cli/v2"
	"os"
	"strings"
)

const (
	Version     = "v0.2.0"
	ProgramName = "btckeys"
)

var (
	BuildTime = "2024-02-19"
	GitCommit = ""
)

const (
	CMD_NAME_GEN = "gen"
)

const (
	CMD_FLAG_NAME_PASSWORD = "password"
	CMD_FLAG_NAME_INDEX    = "index"
	CMD_FLAG_NAME_MNEMONIC = "mnemonic"
	CMD_FLAG_NAME_BIT_SIZE = "bit-size"
)

func init() {
	log.SetLevel("info")
}

func main() {

	local := []*cli.Command{
		genCmd,
	}
	app := &cli.App{
		Name:     ProgramName,
		Version:  fmt.Sprintf("%s %s commit %s", Version, BuildTime, GitCommit),
		Flags:    []cli.Flag{},
		Commands: local,
		Action:   nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

var genCmd = &cli.Command{
	Name:  CMD_NAME_GEN,
	Usage: "generate bitcoin private key",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_PASSWORD,
			Aliases: []string{"p"},
			Usage:   "protect bip39 mnemonic with a passphrase",
		},
		&cli.IntFlag{
			Name:    CMD_FLAG_NAME_INDEX,
			Aliases: []string{"i"},
			Usage:   "index of keys to generate",
			Value:   0,
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_MNEMONIC,
			Aliases: []string{"m"},
			Usage:   "mnemonic words to recover a root key",
		},
		&cli.IntFlag{
			Name:    CMD_FLAG_NAME_BIT_SIZE,
			Aliases: []string{"b"},
			Usage:   "bit size (128/256)",
			Value:   128,
		},
	},
	Action: func(cctx *cli.Context) error {
		password := cctx.String(CMD_FLAG_NAME_PASSWORD)
		index := cctx.Int(CMD_FLAG_NAME_INDEX)
		mnemonic := cctx.String(CMD_FLAG_NAME_MNEMONIC)
		bitsize := cctx.Int(CMD_FLAG_NAME_BIT_SIZE)

		bk, err := btckeys.GenBitcoinKey(password, mnemonic, bitsize, index)
		if err != nil {
			return log.Error(err.Error())
		}
		fmt.Printf(strings.Repeat("-", 149))
		fmt.Printf("\n%-18s %s\n", "BIP39 Mnemonic:", bk.GetMnemonic())
		fmt.Printf("%-18s %s\n", "BIP39 Passphrase:", bk.GetPassphrase())
		fmt.Printf("%-18s %x\n", "BIP39 Seed:", bk.GetSeed())
		fmt.Printf("%-18s %s\n", "BIP32 Root Key:", bk.MasterKeyB58())
		fmt.Printf("\n%-18s %-34s %-42s %-52s\n", "Path(BIP44)", "Bitcoin Address", "Bech32 Address", "WIF(Wallet Import Format)")
		fmt.Printf("%-18s %-34s %s %s\n", bk.GetPath(), bk.Address(), bk.Bech32(), bk.WIF())
		fmt.Printf("\n%-18s %-64s\n", "Path(BIP44)", "Private Key (HEX)")
		fmt.Printf("%-18s %-64s\n", bk.GetPath(), hex.EncodeToString(bk.PrivateKeyBytes()))
		fmt.Printf("\n%-18s %-64s\n", "Path(BIP44)", "Public Key (HEX)")
		fmt.Printf("%-18s %-66s\n", bk.GetPath(), hex.EncodeToString(bk.PublicKeyBytes()))
		fmt.Printf(strings.Repeat("-", 149))
		fmt.Println()
		return nil
	},
}
