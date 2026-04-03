package vcfactory

import (
	"context"
	"fmt"
	"math/big"

	"github.com/HARA-DID/hara-core-blockchain-lib/utils"
)

func (vcf *VCFactory) GetIdentityNFT(ctx context.Context) (utils.Address, error) {
	result, err := vcf.blockchain.CallContract(
		ctx,
		vcf.Contract,
		"getIdentityNFT",
		[]any{},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := vcf.ContractABI.Methods["getIdentityNFT"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}

func (vcf *VCFactory) GetCertificateNFT(ctx context.Context) (utils.Address, error) {
	result, err := vcf.blockchain.CallContract(
		ctx,
		vcf.Contract,
		"getCertificateNFT",
		[]any{},
	)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to call contract: %w", err)
	}

	method := vcf.ContractABI.Methods["getCertificateNFT"]
	unpacked, err := method.Outputs.Unpack(result)
	if err != nil {
		return utils.Address{}, fmt.Errorf("failed to unpack result: %w", err)
	}

	return unpacked[0].(utils.Address), nil
}

func (vcf *VCFactory) GetIdentityTokenIds(
	ctx context.Context,
	did [32]byte,
	offset *big.Int,
	limit *big.Int,
) (*TokenIdsResult, error) {
	result, err := vcf.blockchain.CallContract(
		ctx,
		vcf.Contract,
		"getIdentityTokenIds",
		[]any{did, offset, limit},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var tokenResult TokenIdsResult
	err = vcf.ContractABI.UnpackIntoInterface(&tokenResult, "getIdentityTokenIds", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return &tokenResult, nil
}

func (vcf *VCFactory) GetCertificateTokenIds(
	ctx context.Context,
	did [32]byte,
	offset *big.Int,
	limit *big.Int,
) (*TokenIdsResult, error) {
	result, err := vcf.blockchain.CallContract(
		ctx,
		vcf.Contract,
		"getCertificateTokenIds",
		[]any{did, offset, limit},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var tokenResult TokenIdsResult
	err = vcf.ContractABI.UnpackIntoInterface(&tokenResult, "getCertificateTokenIds", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return &tokenResult, nil
}
