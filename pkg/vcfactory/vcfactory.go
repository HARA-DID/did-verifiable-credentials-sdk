package vcfactory

import (
	"context"
	"fmt"
	"math/big"

	internal "github.com/meQlause/did-verifiable-credentials-sdk/utils"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/contract"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type VCFactory struct {
	blockchain  *blockchain.Blockchain
	ContractABI utils.ABI
	Contract    *contract.Contract
	Address     utils.Address
}

func NewVCFactory(
	contractAddress utils.Address,
	contractABI utils.ABI,
	bc *blockchain.Blockchain,
	contract *contract.Contract,
) *VCFactory {
	return &VCFactory{
		blockchain:  bc,
		ContractABI: contractABI,
		Contract:    contract,
		Address:     contractAddress,
	}
}

func NewVCFactoryWithHNS(
	ctx context.Context,
	hnsURI string,
	bc *blockchain.Blockchain,
) (*VCFactory, error) {
	contract, err := bc.ContractWithHNS(ctx, hnsURI)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve contract with HNS: %w", err)
	}

	return &VCFactory{
		blockchain:  bc,
		Contract:    contract,
		ContractABI: contract.ABI,
		Address:     contract.Address,
	}, nil
}

func (vcf *VCFactory) buildAndSendTx(
	ctx context.Context,
	wallet *wallet.Wallet,
	methodName string,
	params internal.TxParams,
	multipleRPCCalls bool,
) ([]string, error) {
	method, ok := vcf.ContractABI.Methods[methodName]
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

	nonce, err := vcf.blockchain.Network.PendingNonce(ctx, sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	txParams := utils.TransactionParams{
		Nonce:    nonce,
		To:       vcf.Address,
		Value:    big.NewInt(0),
		GasLimit: 30000000,
		GasPrice: big.NewInt(0),
		Data:     calldata,
	}

	tx := vcf.blockchain.BuildTx(txParams)

	hashes, err := vcf.blockchain.CallContractWrite(ctx, wallet, tx, multipleRPCCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return hashes, nil
}

func (vcf *VCFactory) GetAddress() utils.Address {
	return vcf.Address
}

func (vcf *VCFactory) GetBlockchain() *blockchain.Blockchain {
	return vcf.blockchain
}
