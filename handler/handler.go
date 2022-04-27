package handler

import (
	"context"
	"log"
	"math/big"
)

type Monitor interface {
	Run()
	Cancel()
}

type Handler struct {
	txHandler TxHandler
	heightScanner HeightScanner
}

// tx handler
type TxHandler interface {
	HeightScanner
	Do(ctx context.Context, info *TxInfo)
	ListenContract(ctx context.Context, address ContractAddress) bool
}

// block handler
type HeightScanner interface {
	// SaveHeight 持久化最新块高
	SaveHeight(ctx context.Context, height *BlockHeight) error
	// LoadLastHeight 加载上一次块高
	LoadLastHeight(ctx context.Context) (*BlockHeight, error)
}




// 记录高度到 DB
func (block *Handler) SaveHeight(ctx context.Context, height *BlockHeight) error {
	log.Println(height.String())
	return nil
}


func (block *Handler) LoadLastHeight(ctx context.Context) (*BlockHeight, error) {
	return big.NewInt(1), nil
}



