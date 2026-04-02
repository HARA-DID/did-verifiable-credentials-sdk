package vcstorage

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type TokenIdsResult struct {
	TokenIds []*big.Int
	Total    *big.Int
}

type SetAddressParams struct {
	Address utils.Address
}

func (p SetAddressParams) ToArgs() []any {
	return []any{
		p.Address,
	}
}
