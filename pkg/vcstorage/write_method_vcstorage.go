package vcstorage

import (
	"context"

	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/wallet"
)

func (vcs *VCStorage) SetDidRootStorage(
	ctx context.Context,
	wallet *wallet.Wallet,
	params SetAddressParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcs.buildAndSendTx(
		ctx,
		wallet,
		"setDidRootStorage",
		params,
		multipleRPCCalls,
	)
}

func (vcs *VCStorage) SetDidOrgStorage(
	ctx context.Context,
	wallet *wallet.Wallet,
	params SetAddressParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return vcs.buildAndSendTx(
		ctx,
		wallet,
		"setDidOrgStorage",
		params,
		multipleRPCCalls,
	)
}
