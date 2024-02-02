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
	Version     = "v0.1.0"
	ProgramName = "btckeys"
)

var (
	BuildTime = "2024-02-01"
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

		km, err := btckeys.NewKeyManager(bitsize, password, mnemonic)
		if err != nil {
			return log.Errorf(err.Error())
		}
		masterKey, err := km.GetMasterKey()
		if err != nil {
			return log.Errorf(err.Error())
		}
		passphrase := km.GetPassphrase()
		if passphrase == "" {
			passphrase = "<none>"
		}
		log.Printf(strings.Repeat("-", 149))
		log.Printf("%-18s %s", "BIP39 Mnemonic:", km.GetMnemonic())
		log.Printf("%-18s %s", "BIP39 Passphrase:", passphrase)
		log.Printf("%-18s %x", "BIP39 Seed:", km.GetSeed())
		log.Printf("%-18s %s", "BIP32 Root Key:", masterKey.B58Serialize())

		log.Printf("\n%-18s %-34s %-42s %-52s", "Path(BIP44)", "Bitcoin Address", "Bech32 Address", "WIF(Wallet Import Format)")
		key, err := km.GenerateKey(btckeys.PurposeBIP44, btckeys.CoinTypeBTC, 0, 0, uint32(index))
		if err != nil {
			return log.Errorf(err.Error())
		}

		wif, address, segwitBech32, _, err := key.Encode(true)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Printf("%-18s %-34s %s %s", key.GetPath(), address, segwitBech32, wif)

		strPrivateKey := hex.EncodeToString(key.PrivateKey().Serialize())
		strPublicKey := hex.EncodeToString(key.PublicKey().SerializeCompressed())
		log.Printf("\n%-18s %-64s", "Path(BIP44)", "Private Key (HEX)")
		log.Printf("%-18s %-64s", key.GetPath(), strPrivateKey)
		log.Printf("\n%-18s %-64s", "Path(BIP44)", "Public Key (HEX)")
		log.Printf("%-18s %-66s", key.GetPath(), strPublicKey)
		log.Printf(strings.Repeat("-", 149))
		return nil
	},
}
