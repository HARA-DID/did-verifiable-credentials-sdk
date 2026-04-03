package nftbase

import (
	"context"

	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/wallet"
)

func (nft *NFTBase) TransferFrom(
	ctx context.Context,
	wallet *wallet.Wallet,
	params TransferFromParams,
	multipleRPCCalls bool,
) ([]string, error) {
	return nft.buildAndSendTx(
		ctx,
		wallet,
		"transferFrom",
		params,
		multipleRPCCalls,
	)
}
