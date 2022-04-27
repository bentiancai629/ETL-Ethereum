package handler

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"sync"
)

type Monitor interface {
	Run()
	Cancel()
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

var FieldTag = "monitor"

type Handler struct {
	txHandler     TxHandler
	heightScanner HeightScanner
}

type Options struct {
	RpcUrl  string    // eth节点url
	AbiStr  string    // 组合abi,可以监控多个智能合约
	Handler TxHandler // 业务处理handler
	Logger  logrus.FieldLogger
}

type monitor struct {
	ctx     context.Context
	cancel  context.CancelFunc
	cli     *ethclient.Client
	hScan   HeightScanner
	decoder *abiDecoder
	handler TxHandler
	logger  logrus.FieldLogger
	sync.RWMutex
}

func (opt *Options) check() error {
	if opt == nil {
		return errors.New("options nil reference")
	}
	if opt.RpcUrl == "" || opt.AbiStr == "" {
		return errors.New("rpcUrl and abiStr can't be empty")
	}
	if opt.Handler == nil {
		return errors.New("handler nil reference")
	}
	if opt.Logger == nil {
		opt.Logger = logrus.New()
	}
	return nil
}

// 记录高度到 DB
func (block *Handler) SaveHeight(ctx context.Context, height *BlockHeight) error {
	log.Println(height.String())
	return nil
}

func (block *Handler) LoadLastHeight(ctx context.Context) (*BlockHeight, error) {
	return big.NewInt(1), nil
}

