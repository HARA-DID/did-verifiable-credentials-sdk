package vcstorage

import (
	"context"
	"fmt"
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/contract"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	internal "github.com/meQlause/did-verifiable-credentials-sdk/utils"
)

type VCStorage struct {
	blockchain  *blockchain.Blockchain
	ContractABI utils.ABI
	Contract    *contract.Contract
	Address     utils.Address
}

func NewVCStorage(
	contractAddress utils.Address,
	contractABI utils.ABI,
	bc *blockchain.Blockchain,
	contract *contract.Contract,
) *VCStorage {
	return &VCStorage{
		blockchain:  bc,
		ContractABI: contractABI,
		Contract:    contract,
		Address:     contractAddress,
	}
}

func NewVCStorageWithHNS(
	ctx context.Context,
	hnsURI string,
	bc *blockchain.Blockchain,
) (*VCStorage, error) {
	contract, err := bc.ContractWithHNS(ctx, hnsURI)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve contract with HNS: %w", err)
	}

	return &VCStorage{
		blockchain:  bc,
		Contract:    contract,
		ContractABI: contract.ABI,
		Address:     contract.Address,
	}, nil
}

func (vcs *VCStorage) GetAddress() utils.Address {
	return vcs.Address
}

func (vcs *VCStorage) GetBlockchain() *blockchain.Blockchain {
	return vcs.blockchain
}

func (vcs *VCStorage) buildAndSendTx(
	ctx context.Context,
	wallet *wallet.Wallet,
	methodName string,
	params internal.TxParams,
	multipleRPCCalls bool,
) ([]string, error) {
	method, ok := vcs.ContractABI.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("method %s not found in ABI", methodName)
	}

	inputs, err := method.Inputs.Pack(params.ToArgs()...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s arguments: %w", methodName, err)
	}

	calldata := append(method.ID, inputs...)

	sender, err := wallet.GetAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet address: %w", err)
	}

	nonce, err := vcs.blockchain.Network.PendingNonce(ctx, sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	txParams := utils.TransactionParams{
		Nonce:    nonce,
		To:       vcs.Address,
		Value:    big.NewInt(0),
		GasLimit: 3000000,
		GasPrice: big.NewInt(0),
		Data:     calldata,
	}

	tx := vcs.blockchain.BuildTx(txParams)

	hashes, err := vcs.blockchain.CallContractWrite(ctx, wallet, tx, multipleRPCCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return hashes, nil
}
