package types

type BitcoinKey struct {
	*Key
	*KeyManager
}

func (k *BitcoinKey) MasterKeyB58() string {
	mk, _ := k.KeyManager.GetMasterKey()
	return mk.B58Serialize()
}

func (k *BitcoinKey) Address() string {
	_, address, _, _, _ := k.Key.Encode(true)
	return address
}

func (k *BitcoinKey) WIF() string {
	wif, _, _, _, _ := k.Key.Encode(true)
	return wif
}

func (k *BitcoinKey) Bech32() string {
	_, _, segwitBech32, _, _ := k.Key.Encode(true)
	return segwitBech32
}
