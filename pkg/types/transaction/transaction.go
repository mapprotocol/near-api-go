package transaction

import (
	"github.com/eteu-technologies/borsh-go"
	"github.com/mapprotocol/near-api-go/pkg/types/signature"

	"github.com/mapprotocol/near-api-go/pkg/types"
	"github.com/mapprotocol/near-api-go/pkg/types/action"
	"github.com/mapprotocol/near-api-go/pkg/types/hash"
	"github.com/mapprotocol/near-api-go/pkg/types/key"
)

type Transaction struct {
	SignerID   types.AccountID
	PublicKey  key.PublicKey
	Nonce      types.Nonce
	ReceiverID types.AccountID
	BlockHash  hash.CryptoHash
	Actions    []action.Action
}

func (t Transaction) Hash() (txnHash hash.CryptoHash, serialized []byte, err error) {
	// Serialize into Borsh
	serialized, err = borsh.Serialize(t)
	if err != nil {
		return
	}
	txnHash = hash.NewCryptoHash(serialized)
	return
}

func (t Transaction) HashAndSign(keyPair key.KeyPair) (txnHash hash.CryptoHash, serialized []byte, sig signature.Signature, err error) {
	txnHash, serialized, err = t.Hash()
	if err != nil {
		return
	}

	sig = keyPair.Sign(txnHash[:])
	return
}
