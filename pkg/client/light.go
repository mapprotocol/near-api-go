package client

import (
	"context"
	"errors"

	"github.com/mapprotocol/near-api-go/pkg/types/hash"
)

// NextLightClientBlock doc: https://nomicon.io/ChainSpec/LightClient#light-client-block
func (c *Client) NextLightClientBlock(ctx context.Context, lastHash hash.CryptoHash) (resp LightClientBlockView, err error) {
	_, err = c.doRPC(ctx, &resp, "next_light_client_block", nil, []hash.CryptoHash{lastHash})
	return
}

// LightClientProof doc: https://nomicon.io/ChainSpec/LightClient#light-client-proof
func (c *Client) LightClientProof(ctx context.Context, tor TransactionOrReceiptId) (resp RpcLightClientExecutionProofResponse, err error) {
	params := make(map[string]interface{}, 3)
	typ := tor.Type()
	if typ == TypeTransaction {
		params["transaction_hash"] = tor.Hash()
		params["sender_id"] = tor.ID()
	} else if typ == TypeReceipt {
		params["receipt_id"] = tor.Hash()
		params["receiver_id"] = tor.ID()
	} else {
		return RpcLightClientExecutionProofResponse{}, errors.New("invalid type")
	}
	params["type"] = typ
	params["light_client_head"] = tor.LightClientHeadHash()

	_, err = c.doRPC(ctx, &resp, "EXPERIMENTAL_light_client_proof", nil, params)
	return
}
