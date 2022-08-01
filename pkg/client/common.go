package client

import (
	"github.com/mapprotocol/near-api-go/pkg/types"
	"github.com/mapprotocol/near-api-go/pkg/types/hash"
)

type QueryResponse struct {
	BlockHeight types.BlockHeight `json:"block_height"`
	BlockHash   hash.CryptoHash   `json:"block_hash"`
	Error       *string           `json:"error"`
	Logs        []interface{}     `json:"logs"` // TODO: use correct type
}
