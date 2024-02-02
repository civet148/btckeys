# btckeys
Bitcoin keys generation and recover

# build

```shell
$ make
rm -f btckeys
go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=`date +'%Y%m%d %H:%M:%S'`' -X 'main.GitCommit=`git rev-parse --short HEAD`'" -o btckeys cmd/main.go
```

# help

```shell
$ ./btckeys gen -h
NAME:
   btckeys gen - generate bitcoin private key

USAGE:
   btckeys gen [command options] [arguments...]

OPTIONS:
   --password value, -p value  protect bip39 mnemonic with a passphrase
   --index  value, -i value    index of keys to generate (default: 0)
   --mnemonic value, -m value  mnemonic words to recover a root key
   --bit-size value, -b value  bit size (128/256) (default: 128)
   --help, -h                  show help (default: false)
```

# usage

- generate key 

```shell
$ make && ./btckeys gen 
-----------------------------------------------------------------------------------------------------------------------------------------------------
BIP39 Mnemonic:    portion regular fly hip jaguar limit cabin this use normal tiger disagree
BIP39 Passphrase:  <none>
BIP39 Seed:        06e663f7077761c2a01c1cdd97c0c18dbbf8f23489bc7ea0a9ed7eb317c17a7a4ea1a530ec9885ef43a782a1f8179da41a33614ca7f9c4bb83b566061af2b6be
BIP32 Root Key:    xprv9s21ZrQH143K2p1BVs4raMRC3soMJnNarypNzkFuPgxPRtWjpFqpS8ZXnhjxh7rDbB8QxDn4ajaLkbT3PHNLE2pgZZC545McboPAERTMom6

Path(BIP44)        Bitcoin Address                    Bech32 Address                             WIF(Wallet Import Format)
m/44'/0'/0'/0/0    18P3kFeR2ASBLwxC6QMz6GB7cqgJP8E9zC bc1q2r699klfh0wl9k3v80k4myny2cw76f7qq49lpg L4fdzfMBWehaMyaqnPfrAsKTvfwAVTCEnTnQyGTSwkNahLLDe1F9

Path(BIP44)        Private Key (HEX)
m/44'/0'/0'/0/0    de3333408d5d987adfef51e3090ff8cfba28faa84167c7bdfd07a926ce0b3b89

Path(BIP44)        Public Key (HEX)
m/44'/0'/0'/0/0    038fd6642174d29fda7c14e3c23e892690b2be21d9396ee7ab39263b5ca610529e
-----------------------------------------------------------------------------------------------------------------------------------------------------
```

- re-generate by mnemonic
```shell
$ make && ./btckeys gen --mnemonic "move potato author cherry strong cook near task expect town relief fuel" 

-----------------------------------------------------------------------------------------------------------------------------------------------------
BIP39 Mnemonic:    move potato author cherry strong cook near task expect town relief fuel
BIP39 Passphrase:  <none>
BIP39 Seed:        3893a65e7955fe0b212fdc672d0272df29b0c05cbd6a3057f4a4e678e1084459027202a2e02c9032260f64721cec7dec91f7b41cda00306fd1bd237ad206d9ad
BIP32 Root Key:    xprv9s21ZrQH143K2MBdrkX55VNXagjbMuYtHBRFZzwiJYxGuhpZHLq94ZYLeycKSJSGm7AjMFxqnb8d5dySADoLtW41XDxpCKZzvwcjhttWU1G

Path(BIP44)        Bitcoin Address                    Bech32 Address                             WIF(Wallet Import Format)
m/44'/0'/0'/0/0    1CtHLL4yabJMpoUeJ7aavJVGi5AEUkThxX bc1qsfwrkcuzde4er9vt539ntg9rnkxrcvnpt0pcqr L2CtraBeSGXzyYusi9rXhRXJXCmhEqtzAdYoSVASKV7M8Rpe5AMr

Path(BIP44)        Private Key (HEX)
m/44'/0'/0'/0/0    94c466b73ac761e0be2027e7b734a6ddc3c64e5f7bd4d430a477509aa3cc008e

Path(BIP44)        Public Key (HEX)
m/44'/0'/0'/0/0    030a65ebb842e58a2a8edc117e700bdb3afcc79a3555c9017221b172bfae375618
-----------------------------------------------------------------------------------------------------------------------------------------------------
```