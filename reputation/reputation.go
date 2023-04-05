// Package reputation implements contract function
package reputation

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	reputation "github.com/ethereum/go-ethereum/reputation/contract"
)

// https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings
//go:generate solc @openzeppelin=../node_modules/@openzeppelin ./contract/reputation.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ./contract/ --overwrite
//go:generate go run ../cmd/abigen --pkg reputation --out ./contract/reputation_helper.go --combined-json ./contract/combined.json

var ReputationAccount = "0x3967df487B277e67E5B4Ca5397C6F6C7875B64B4"
var ReputationABI = reputation.RepValueMetaData.ABI
var ReputationBin = reputation.RepValueMetaData.Bin

type Reputation struct {
	log log.Logger

	contract *reputation.RepValue
	caller   *bind.TransactOpts
}

func New(log log.Logger) *Reputation {
	return &Reputation{log: log}
}

func DeployContract(client *ethclient.Client, privateKeyAddress string) (common.Address, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress=", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000000)

	//deploying smart contract
	address, tx, instance, err := reputation.DeployRepValue(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
	fmt.Println("instance->", instance)
	fmt.Println("tx->", tx.Hash().Hex())

	return address, nil
}

// BindContract 绑定已经部署在区块链上的合约
func (rep *Reputation) BindContract(client *ethclient.Client, privateKeyAddress string, contractAddress common.Address) error {

	caller, err := rep.getCaller(client, privateKeyAddress)
	if err != nil {
		rep.log.Error("BindContract", "could not get caller may not update reputation", err)
	}

	rep.caller = caller

	contract, err := reputation.NewRepValue(contractAddress, client)
	if err != nil {
		rep.log.Error("BindContract", "bind contract error", err)
		return err
	}

	rep.contract = contract

	return nil
}

// GetReputation 获取信誉
func (rep *Reputation) GetReputation(account common.Address) (*big.Int, error) {
	return rep.contract.GetReputation(nil, account)
}

// GetRepThreshold 获取阈值
func (rep *Reputation) GetRepThreshold() (*big.Int, error) {
	return rep.contract.GetRepThreshold(nil)
}

// UpdateReputation 更新信誉
func (rep *Reputation) UpdateReputation(account common.Address, score *big.Int) (*types.Transaction, error) {
	return rep.contract.UpdateReputation(rep.caller, account, score)
}

func (rep *Reputation) getCaller(client *ethclient.Client, privateKeyAddress string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		rep.log.Warn("getCaller", "HexToECDSA", err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		rep.log.Warn("getCaller", "invalid key", err)
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	rep.log.Info("getCaller", "fromAddress = ", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		rep.log.Warn("getCaller", "PendingNonceAt", err)
		return nil, err
	}
	rep.log.Info("getCaller", "nonce", nonce)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		rep.log.Warn("getCaller", "ChainID", err)
		return nil, err
	}

	caller, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		rep.log.Warn("getCaller", "NewKeyedTransactorWithChainID", err)
		return nil, err
	}

	caller.Nonce = big.NewInt(int64(nonce))
	caller.Value = big.NewInt(0)      // in wei
	caller.GasLimit = uint64(3000000) // in units
	caller.GasPrice = big.NewInt(1000000000)

	return caller, nil
}
