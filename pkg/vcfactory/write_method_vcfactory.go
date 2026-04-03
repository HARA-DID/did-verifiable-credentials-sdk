package vcfactory

import (
	"context"

	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/wallet"
)

func (vcf *VCFactory) IssueCredential(
	ctx context.Context,
	wallet *wallet.Wallet,
	params IssueCredentialParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcf.buildAndSendTx(
		ctx,
		wallet,
		"issueCredential",
		params,
		multipleRPCCalls,
	)
}

func (vcf *VCFactory) BurnCredential(
	ctx context.Context,
	wallet *wallet.Wallet,
	params BurnCredentialParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcf.buildAndSendTx(
		ctx,
		wallet,
		"burnCredential",
		params,
		multipleRPCCalls,
	)
}

func (vcf *VCFactory) UpdateMetadata(
	ctx context.Context,
	wallet *wallet.Wallet,
	params UpdateMetadataParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcf.buildAndSendTx(
		ctx,
		wallet,
		"updateMetadata",
		params,
		multipleRPCCalls,
	)
}

func (vcf *VCFactory) RevokeCredential(
	ctx context.Context,
	wallet *wallet.Wallet,
	params RevokeCredentialParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcf.buildAndSendTx(
		ctx,
		wallet,
		"revokeCredential",
		params,
		multipleRPCCalls,
	)
}

func (vcf *VCFactory) ApproveCredentialOrg(
	ctx context.Context,
	wallet *wallet.Wallet,
	params ApproveCredentialOrgParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcf.buildAndSendTx(
		ctx,
		wallet,
		"approveCredentialOrg",
		params,
		multipleRPCCalls,
	)
}

func (vcf *VCFactory) ApproveCredential(
	ctx context.Context,
	wallet *wallet.Wallet,
	params ApproveCredentialParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcf.buildAndSendTx(
		ctx,
		wallet,
		"approveCredential",
		params,
		multipleRPCCalls,
	)
}
