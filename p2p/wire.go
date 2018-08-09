package p2p

import (
	"github.com/tendermint/go-amino"
	"github.com/9thchain/blockchain-p2p/crypto"
)

var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}
