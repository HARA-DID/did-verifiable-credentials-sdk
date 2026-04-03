package nftbase

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	internal "github.com/HARA-DID/did-verifiable-credentials-sdk/utils"

	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/contract"
	"github.com/HARA-DID/hara-core-blockchain-lib/pkg/wallet"
	"github.com/HARA-DID/hara-core-blockchain-lib/utils"
)

type NFTBase struct {
	blockchain  *blockchain.Blockchain
	ContractABI utils.ABI
	Contract    *contract.Contract
	Address     utils.Address
}

func NewNFTBase(
	contractAddress utils.Address,
	contractABI utils.ABI,
	bc *blockchain.Blockchain,
	contract *contract.Contract,
) *NFTBase {
	return &NFTBase{
		blockchain:  bc,
		ContractABI: contractABI,
		Contract:    contract,
		Address:     contractAddress,
	}
}

func NewNFTBaseWithHNS(
	ctx context.Context,
	hnsURI string,
	bc *blockchain.Blockchain,
) (*NFTBase, error) {
	contract, err := bc.ContractWithHNS(ctx, hnsURI)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve contract with HNS: %w", err)
	}

	return &NFTBase{
		blockchain:  bc,
		Contract:    contract,
		ContractABI: contract.ABI,
		Address:     contract.Address,
	}, nil
}

func (nft *NFTBase) GetAddress() utils.Address {
	return nft.Address
}

func (nft *NFTBase) GetBlockchain() *blockchain.Blockchain {
	return nft.blockchain
}

func (nft *NFTBase) call(ctx context.Context, method string, args ...any) ([]byte, error) {
	data, err := nft.ContractABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("abi pack error for %s: %w", method, err)
	}
	raw := "0x" + utils.Bytes2Hex(data)
	return nft.blockchain.Network.Call(ctx, nft.Address, raw)
}

func (nft *NFTBase) buildAndSendTx(
	ctx context.Context,
	wallet *wallet.Wallet,
	methodName string,
	params internal.TxParams,
	multipleRPCCalls bool,
) ([]string, error) {
	method, ok := nft.ContractABI.Methods[methodName]
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

	nonce, err := nft.blockchain.Network.PendingNonce(ctx, sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	txParams := utils.TransactionParams{
		Nonce:    nonce,
		To:       nft.Address,
		Value:    big.NewInt(0),
		GasLimit: 30000000,
		GasPrice: big.NewInt(0),
		Data:     calldata,
	}

	tx := nft.blockchain.BuildTx(txParams)

	hashes, err := nft.blockchain.CallContractWrite(ctx, wallet, tx, multipleRPCCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return hashes, nil
}

func unwrapDoubleEncoding(out []byte) ([]byte, error) {
	if len(out) > 2 && out[0] == 0x22 && out[len(out)-1] == 0x22 {
		asciiStr := string(out)
		innerHex := strings.Trim(asciiStr, "\"")
		innerBytes, err := hex.DecodeString(strings.TrimPrefix(innerHex, "0x"))
		if err != nil {
			return nil, fmt.Errorf("failed to decode inner hex: %w", err)
		}
		return innerBytes, nil
	}
	return out, nil
}
