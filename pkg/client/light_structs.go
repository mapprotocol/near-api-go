package client

import (
	"github.com/mapprotocol/near-api-go/pkg/types"
	"github.com/mapprotocol/near-api-go/pkg/types/hash"
	"github.com/mapprotocol/near-api-go/pkg/types/key"
	"github.com/mapprotocol/near-api-go/pkg/types/signature"
)

type ValidatorStakeWithVersion struct {
	AccountID                   types.AccountID     `json:"account_id"`
	PublicKey                   key.Base58PublicKey `json:"public_key"`
	Stake                       types.Balance       `json:"stake"`
	ValidatorStakeStructVersion string              `json:"validator_stake_struct_version"`
}

type BlockHeaderInnerLiteView struct {
	Height           types.BlockHeight `json:"height"`
	EpochID          hash.CryptoHash   `json:"epoch_id"`
	NextEpochId      hash.CryptoHash   `json:"next_epoch_id"`
	PrevStateRoot    hash.CryptoHash   `json:"prev_state_root"`
	OutcomeRoot      hash.CryptoHash   `json:"outcome_root"`
	Timestamp        uint64            `json:"timestamp"`
	TimestampNanoSec string            `json:"timestamp_nanosec"`
	NextBpHash       hash.CryptoHash   `json:"next_bp_hash"`
	BlockMerkleRoot  hash.CryptoHash   `json:"block_merkle_root"`
}

type LightClientBlockView struct {
	PrevBlockHash      hash.CryptoHash              `json:"prev_block_hash"`
	NextBlockInnerHash hash.CryptoHash              `json:"next_block_inner_hash"`
	InnerLite          BlockHeaderInnerLiteView     `json:"inner_lite"`
	InnerRestHash      hash.CryptoHash              `json:"inner_rest_hash"`
	NextBps            []ValidatorStakeWithVersion  `json:"next_bps"`
	ApprovalsAfterNext []*signature.Base58Signature `json:"approvals_after_next"`
}

type LightClientBlockLiteView struct {
	PrevBlockHash hash.CryptoHash          `json:"prev_block_hash"`
	InnerRestHash hash.CryptoHash          `json:"inner_rest_hash"`
	InnerLite     BlockHeaderInnerLiteView `json:"inner_lite"`
}

type RpcLightClientExecutionProofResponse struct {
	OutcomeProof     ExecutionOutcomeWithIdView `json:"outcome_proof"`
	OutcomeRootProof MerklePath                 `json:"outcome_root_proof"`
	BlockHeaderLite  LightClientBlockLiteView   `json:"block_header_lite"`
	BlockProof       MerklePath                 `json:"block_proof"`
}

const (
	TypeReceipt     = "receipt"
	TypeTransaction = "transaction"
)

type TransactionOrReceiptId interface {
	Type() string
	ID() types.AccountID
	Hash() hash.CryptoHash
	LightClientHeadHash() hash.CryptoHash
}

type Transaction struct {
	TransactionHash hash.CryptoHash `json:"transaction_hash"`
	SenderID        types.AccountID `json:"sender_id"`
	LightClientHead hash.CryptoHash `json:"light_client_head"`
}

func (t Transaction) Type() string {
	return TypeTransaction
}

func (t Transaction) ID() types.AccountID {
	return t.SenderID
}

func (t Transaction) Hash() hash.CryptoHash {
	return t.TransactionHash
}

func (t Transaction) LightClientHeadHash() hash.CryptoHash {
	return t.LightClientHead
}

type Receipt struct {
	ReceiptID       hash.CryptoHash `json:"receipt_id"`
	ReceiverID      types.AccountID `json:"receiver_id"`
	LightClientHead hash.CryptoHash
}

func (r Receipt) Type() string {
	return TypeReceipt
}

func (r Receipt) ID() types.AccountID {
	return r.ReceiverID
}

func (r Receipt) Hash() hash.CryptoHash {
	return r.ReceiptID
}

func (r Receipt) LightClientHeadHash() hash.CryptoHash {
	return r.LightClientHead
}
