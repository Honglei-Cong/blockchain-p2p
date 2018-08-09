package conn

import (
	"github.com/tendermint/go-amino"
	"github.com/9thchain/blockchain-p2p/crypto"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
