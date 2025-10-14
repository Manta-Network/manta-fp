package types

import (
	"github.com/Manta-Network/manta-fp/finality-provider/proto"
	bbntypes "github.com/babylonlabs-io/babylon/v3/types"
	"github.com/btcsuite/btcd/btcec/v2"
)

type FinalityProviderState interface {
	GetBtcPk() *btcec.PublicKey
	GetBtcPkBIP340() *bbntypes.BIP340PubKey
	GetBtcPkHex() string
	GetChainID() []byte
	GetLastVotedHeight() uint64
	SetLastVotedHeight(height uint64) error
	GetStatus() proto.FinalityProviderStatus
	SetStatus(status proto.FinalityProviderStatus) error
}
