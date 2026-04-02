package vcfactory

import "math/big"

type Options uint8

const (
	Identity Options = iota
	Certificate
)

type IssueCredentialParams struct {
	Option         Options
	DIDRecipient   [32]byte
	Issuer         [32]byte
	ExpiredAt      *big.Int
	OffchainHash   [32]byte
	MerkleTreeRoot [32]byte
	PublicIdentity [32]byte
}

func (p IssueCredentialParams) ToArgs() []any {
	return []any{
		uint8(p.Option),
		p.DIDRecipient,
		p.Issuer,
		p.ExpiredAt,
		p.OffchainHash,
		p.MerkleTreeRoot,
		p.PublicIdentity,
	}
}

type BurnCredentialParams struct {
	Option  Options
	DID     [32]byte
	TokenID *big.Int
}

func (p BurnCredentialParams) ToArgs() []any {
	return []any{
		uint8(p.Option),
		p.DID,
		p.TokenID,
	}
}

type UpdateMetadataParams struct {
	Option       Options
	TokenID      *big.Int
	ExpiredAt    *big.Int
	OffchainHash [32]byte
}

func (p UpdateMetadataParams) ToArgs() []any {
	return []any{
		uint8(p.Option),
		p.TokenID,
		p.ExpiredAt,
		p.OffchainHash,
	}
}

type RevokeCredentialParams struct {
	Option  Options
	TokenID *big.Int
}

func (p RevokeCredentialParams) ToArgs() []any {
	return []any{
		uint8(p.Option),
		p.TokenID,
	}
}

type ApproveCredentialOrgParams struct {
	Option      Options
	TokenID     *big.Int
	OrgDIDHash  [32]byte
	UserDIDHash [32]byte
	Signature   []byte
}

func (p ApproveCredentialOrgParams) ToArgs() []any {
	return []any{
		uint8(p.Option),
		p.TokenID,
		p.OrgDIDHash,
		p.UserDIDHash,
		p.Signature,
	}
}

type ApproveCredentialParams struct {
	Option    Options
	TokenID   *big.Int
	Signature []byte
}

func (p ApproveCredentialParams) ToArgs() []any {
	return []any{
		uint8(p.Option),
		p.TokenID,
		p.Signature,
	}
}

type TokenIdsResult struct {
	TokenIds []*big.Int
	Total    *big.Int
}

type CredentialMetadata struct {
	IsValid      bool
	ExpiredAt    *big.Int
	Issuer       [32]byte 
	IssuedAt     *big.Int
	OffchainHash [32]byte
	Claimed      bool
}