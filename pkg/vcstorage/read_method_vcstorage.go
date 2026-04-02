package vcstorage

import (
	"context"
	"fmt"
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (vcs *VCStorage) GetIdentityTokenCount(
	ctx context.Context,
	did [32]byte,
) (*big.Int, error) {
	result, err := vcs.blockchain.CallContract(
		ctx,
		vcs.Contract,
		"getIdentityTokenCount",
		[]any{did},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	method := vcs.ContractABI.Methods["getIdentityTokenCount"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(*big.Int), nil
}

func (vcs *VCStorage) GetCertificateTokenCount(
	ctx context.Context,
	did [32]byte,
) (*big.Int, error) {
	result, err := vcs.blockchain.CallContract(
		ctx,
		vcs.Contract,
		"getCertificateTokenCount",
		[]any{did},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	method := vcs.ContractABI.Methods["getCertificateTokenCount"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(*big.Int), nil
}

func (vcs *VCStorage) GetIdentityTokenIds(
	ctx context.Context,
	did [32]byte,
	offset *big.Int,
	limit *big.Int,
) (*TokenIdsResult, error) {
	result, err := vcs.blockchain.CallContract(
		ctx,
		vcs.Contract,
		"getIdentityTokenIds",
		[]any{did, offset, limit},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var tokenResult TokenIdsResult
	err = vcs.ContractABI.UnpackIntoInterface(&tokenResult, "getIdentityTokenIds", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return &tokenResult, nil
}

func (vcs *VCStorage) GetCertificateTokenIds(
	ctx context.Context,
	did [32]byte,
	offset *big.Int,
	limit *big.Int,
) (*TokenIdsResult, error) {
	result, err := vcs.blockchain.CallContract(
		ctx,
		vcs.Contract,
		"getCertificateTokenIds",
		[]any{did, offset, limit},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var tokenResult TokenIdsResult
	err = vcs.ContractABI.UnpackIntoInterface(&tokenResult, "getCertificateTokenIds", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return &tokenResult, nil
}


func (vcs *VCStorage) GetDIDRootStorage(ctx context.Context) (utils.Address, error) {
	result, err := vcs.blockchain.CallContract(
		ctx,
		vcs.Contract,
		"getDIDRootStorage",
		[]any{},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := vcs.ContractABI.Methods["getDIDRootStorage"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}

func (vcs *VCStorage) GetDIDOrgStorage(ctx context.Context) (utils.Address, error) {
	result, err := vcs.blockchain.CallContract(
		ctx,
		vcs.Contract,
		"getDIDOrgStorage",
		[]any{},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := vcs.ContractABI.Methods["getDIDOrgStorage"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}
