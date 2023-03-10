// Package reputation implements contract function
package reputation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	reputation "github.com/ethereum/go-ethereum/reputation/contract"
)

//go:generate solc @openzeppelin=../node_modules/@openzeppelin contract/reputation.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ./contract/ --overwrite
//go:generate go run ../cmd/abigen --pkg reputation --out ./contract/reputation_helper.go --combined-json ./contract/combined.json

type Reputations struct {
	contract *reputation.Reputation
	caller   *bind.TransactOpts
}

func NewReputations(privateKeyAddress string) *Reputations {
	return &Reputations{}
}

func DeployContract() {
	// 部署合约-账户？

	// 合约地址写入创世块的数据库里
}

// GetReputation 获取信誉
func (r *Reputations) GetReputation(address common.Address) (*big.Int, error) {

	return r.contract.GetReputation(nil, address)
}

// GetRepThreshold 获取阈值
func (r *Reputations) GetRepThreshold() (*big.Int, error) {

	return r.contract.GetRepThreshold(nil)
}

// UpdateReputation 更新信誉
func (r *Reputations) UpdateReputation(address common.Address, score *big.Int) (*types.Transaction, error) {

	return r.contract.UpdateReputation(r.caller, address, score)
}
