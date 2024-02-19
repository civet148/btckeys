package btckeys

import (
	"encoding/hex"
	"fmt"
	"github.com/civet148/log"
	"strings"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	bk, err := GenBitcoinKey("", "", 128, 0)
	if err != nil {
		log.Error(err.Error())
		return
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
}
