package nftbase

import (
	"math/big"

	"github.com/HARA-DID/hara-core-blockchain-lib/utils"
)

type MintParams struct {
	Issuer         utils.Address
	ClaimerDIDHash [32]byte
}

func (p MintParams) ToArgs() []any {
	return []any{
		p.Issuer,
		p.ClaimerDIDHash,
	}
}

type BurnParams struct {
	TokenID *big.Int
}

func (p BurnParams) ToArgs() []any {
	return []any{
		p.TokenID,
	}
}

type SetMetadataParams struct {
	TokenID  *big.Int
	Metadata CredentialMetadata
}

func (p SetMetadataParams) ToArgs() []any {
	return []any{
		p.TokenID,
		p.Metadata,
	}
}

type TransferFromParams struct {
	From    utils.Address
	To      utils.Address
	TokenID *big.Int
}

func (p TransferFromParams) ToArgs() []any {
	return []any{
		p.From,
		p.To,
		p.TokenID,
	}
}

type AddTokenToDIDParams struct {
	DID     [32]byte
	TokenID *big.Int
}

func (p AddTokenToDIDParams) ToArgs() []any {
	return []any{
		p.DID,
		p.TokenID,
	}
}

type RemoveTokenFromDIDParams struct {
	DID     [32]byte
	TokenID *big.Int
}

func (p RemoveTokenFromDIDParams) ToArgs() []any {
	return []any{
		p.DID,
		p.TokenID,
	}
}

type CredentialMetadata struct {
	IsValid      bool
	ExpiredAt    *big.Int
	Issuer       utils.Address
	IssuedAt     *big.Int
	OffchainHash string
	Claimed      bool
}

type CredentialsWithMetadataResult struct {
	TokenIds []*big.Int
	Metadata []CredentialMetadata
}
