package nftbase

import (
	"context"
	"fmt"
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (nft *NFTBase) IsCredentialValid(
	ctx context.Context,
	tokenId *big.Int,
) (bool, error) {
	out, err := nft.call(ctx, "isCredentialValid", tokenId)
	if err != nil {
		return false, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return false, err
	}

	values, err := nft.ContractABI.Methods["isCredentialValid"].Outputs.Unpack(out)
	if err != nil {
		return false, fmt.Errorf("decode isCredentialValid: %w", err)
	}
	if len(values) != 1 {
		return false, fmt.Errorf("unexpected isCredentialValid result length: %d", len(values))
	}

	isValid, ok := values[0].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected isCredentialValid type %T", values[0])
	}

	return isValid, nil
}

func (nft *NFTBase) GetMetadata(
	ctx context.Context,
	tokenId *big.Int,
) (*CredentialMetadata, error) {
	out, err := nft.call(ctx, "getMetadata", tokenId)
	if err != nil {
		return nil, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return nil, err
	}

	values, err := nft.ContractABI.Methods["getMetadata"].Outputs.Unpack(out)
	if err != nil {
		return nil, fmt.Errorf("decode getMetadata: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("unexpected getMetadata result length: %d", len(values))
	}

	// The result should be a struct with the metadata fields - use json tags
	result := values[0].(struct {
		IsValid      bool          `json:"isValid"`
		ExpiredAt    *big.Int      `json:"expiredAt"`
		Issuer       utils.Address `json:"issuer"`
		IssuedAt     *big.Int      `json:"issuedAt"`
		OffchainHash string        `json:"offchainHash"`
		Claimed      bool          `json:"claimed"`
	})

	metadata := &CredentialMetadata{
		IsValid:      result.IsValid,
		ExpiredAt:    result.ExpiredAt,
		Issuer:       result.Issuer,
		IssuedAt:     result.IssuedAt,
		OffchainHash: result.OffchainHash,
		Claimed:      result.Claimed,
	}

	return metadata, nil
}

func (nft *NFTBase) GetCredentialsWithMetadata(
	ctx context.Context,
	tokenIds []*big.Int,
) (*CredentialsWithMetadataResult, error) {
	out, err := nft.call(ctx, "getCredentialsWithMetadata", tokenIds)
	if err != nil {
		return nil, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return nil, err
	}

	values, err := nft.ContractABI.Methods["getCredentialsWithMetadata"].Outputs.Unpack(out)
	if err != nil {
		return nil, fmt.Errorf("decode getCredentialsWithMetadata: %w", err)
	}
	if len(values) != 2 {
		return nil, fmt.Errorf("unexpected getCredentialsWithMetadata result length: %d", len(values))
	}

	// First value is tokenIds array
	returnedTokenIds, ok := values[0].([]*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected tokenIds type %T", values[0])
	}

	// Second value is metadata array - use json tags as that's what ABI unpacker uses
	metadataArray, ok := values[1].([]struct {
		IsValid      bool          `json:"isValid"`
		ExpiredAt    *big.Int      `json:"expiredAt"`
		Issuer       utils.Address `json:"issuer"`
		IssuedAt     *big.Int      `json:"issuedAt"`
		OffchainHash string        `json:"offchainHash"`
		Claimed      bool          `json:"claimed"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected metadata type %T", values[1])
	}

	// Convert to our result struct
	credResult := &CredentialsWithMetadataResult{
		TokenIds: returnedTokenIds,
		Metadata: make([]CredentialMetadata, len(metadataArray)),
	}

	for i, meta := range metadataArray {
		credResult.Metadata[i] = CredentialMetadata{
			IsValid:      meta.IsValid,
			ExpiredAt:    meta.ExpiredAt,
			Issuer:       meta.Issuer,
			IssuedAt:     meta.IssuedAt,
			OffchainHash: meta.OffchainHash,
			Claimed:      meta.Claimed,
		}
	}

	return credResult, nil
}

func (nft *NFTBase) GetUnclaimedTokenId(
	ctx context.Context,
	tokenId *big.Int,
) (utils.Hash, error) {
	out, err := nft.call(ctx, "getUnclaimedTokenId", tokenId)
	if err != nil {
		return utils.Hash{}, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return utils.Hash{}, err
	}

	values, err := nft.ContractABI.Methods["getUnclaimedTokenId"].Outputs.Unpack(out)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("decode getUnclaimedTokenId: %w", err)
	}
	if len(values) != 1 {
		return utils.Hash{}, fmt.Errorf("unexpected getUnclaimedTokenId result length: %d", len(values))
	}

	var hash utils.Hash
	switch v := values[0].(type) {
	case [32]byte:
		hash = utils.Hash(v)
	case utils.Hash:
		hash = v
	default:
		return utils.Hash{}, fmt.Errorf("unexpected getUnclaimedTokenId type %T", values[0])
	}

	return hash, nil
}

func (nft *NFTBase) Exists(
	ctx context.Context,
	tokenId *big.Int,
) (bool, error) {
	out, err := nft.call(ctx, "exists", tokenId)
	if err != nil {
		return false, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return false, err
	}

	values, err := nft.ContractABI.Methods["exists"].Outputs.Unpack(out)
	if err != nil {
		return false, fmt.Errorf("decode exists: %w", err)
	}
	if len(values) != 1 {
		return false, fmt.Errorf("unexpected exists result length: %d", len(values))
	}

	exists, ok := values[0].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected exists type %T", values[0])
	}

	return exists, nil
}

func (nft *NFTBase) OwnerOf(
	ctx context.Context,
	tokenId *big.Int,
) (utils.Address, error) {
	out, err := nft.call(ctx, "ownerOf", tokenId)
	if err != nil {
		return utils.Address{}, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return utils.Address{}, err
	}

	values, err := nft.ContractABI.Methods["ownerOf"].Outputs.Unpack(out)
	if err != nil {
		return utils.Address{}, fmt.Errorf("decode ownerOf: %w", err)
	}
	if len(values) != 1 {
		return utils.Address{}, fmt.Errorf("unexpected ownerOf result length: %d", len(values))
	}

	owner, ok := values[0].(utils.Address)
	if !ok {
		return utils.Address{}, fmt.Errorf("unexpected ownerOf type %T", values[0])
	}

	return owner, nil
}

func (nft *NFTBase) TotalTokensToBeClaimedByDid(
	ctx context.Context,
	did utils.Hash,
) (*big.Int, error) {
	out, err := nft.call(ctx, "totalTokensToBeClaimedByDid", did)
	if err != nil {
		return nil, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return nil, err
	}

	values, err := nft.ContractABI.Methods["totalTokensToBeClaimedByDid"].Outputs.Unpack(out)
	if err != nil {
		return nil, fmt.Errorf("decode totalTokensToBeClaimedByDid: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("unexpected totalTokensToBeClaimedByDid result length: %d", len(values))
	}

	total, ok := values[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected totalTokensToBeClaimedByDid type %T", values[0])
	}

	return total, nil
}

func (nft *NFTBase) GetToBeClaimedTokensByDid(
	ctx context.Context,
	did utils.Hash,
	offset *big.Int,
	limit *big.Int,
) ([]*big.Int, error) {
	out, err := nft.call(ctx, "getToBeClaimedTokensByDid", did, offset, limit)
	if err != nil {
		return nil, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return nil, err
	}

	values, err := nft.ContractABI.Methods["getToBeClaimedTokensByDid"].Outputs.Unpack(out)
	if err != nil {
		return nil, fmt.Errorf("decode getToBeClaimedTokensByDid: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("unexpected getToBeClaimedTokensByDid result length: %d", len(values))
	}

	tokens, ok := values[0].([]*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected getToBeClaimedTokensByDid type %T", values[0])
	}

	return tokens, nil
}

func (nft *NFTBase) IsApprovedForAll(
	ctx context.Context,
	owner utils.Address,
	operator utils.Address,
) (bool, error) {
	out, err := nft.call(ctx, "isApprovedForAll", owner, operator)
	if err != nil {
		return false, err
	}

	out, err = unwrapDoubleEncoding(out)
	if err != nil {
		return false, err
	}

	values, err := nft.ContractABI.Methods["isApprovedForAll"].Outputs.Unpack(out)
	if err != nil {
		return false, fmt.Errorf("decode isApprovedForAll: %w", err)
	}
	if len(values) != 1 {
		return false, fmt.Errorf("unexpected isApprovedForAll result length: %d", len(values))
	}

	isApproved, ok := values[0].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected isApprovedForAll type %T", values[0])
	}

	return isApproved, nil
}
