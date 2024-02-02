package btckeys

import (
	"github.com/civet148/btckeys/types"
	"github.com/civet148/log"
)

func GenerateKey(passphrase, mnemonic string, bitSize, index int) (*types.BitcoinKey, error) {
	km, err := types.NewKeyManager(bitSize, passphrase, mnemonic)
	if err != nil {
		return nil, log.Errorf(err.Error())
	}
	key, err := km.GenerateKey(types.PurposeBIP44, types.CoinTypeBTC, 0, 0, uint32(index))
	if err != nil {
		return nil, log.Errorf(err.Error())
	}
	return &types.BitcoinKey{
		Key:        key,
		KeyManager: km,
	}, nil
}
