// Package repvm 基于clique，实现信誉值管理
package repvm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/core/types"
)

// Repvm 基于clique共识修改，实现信誉值管理
type Repvm struct {
	*clique.Clique // 基于clique的共识
}

// New 初始化repvm共识
func New(origin *clique.Clique) *Repvm {
	// 返回repvm实例
	return &Repvm{origin}
}

func (r *Repvm) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	//TODO implement me
	panic("implement me")
}

func (r Repvm) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	//TODO implement me
	panic("implement me")
}
