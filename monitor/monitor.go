package monitor

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
	"time"
)

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

func (m *monitor) Run() {
	lastBlockHeight, err := m.hScan.LoadLastHeight(m.ctx)
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			curIndex, _, err := m.getBlockHeight()
			if err != nil {
				m.logger.WithField(FieldTag, "getBlockHeight").Error(err)
				continue
			}
			if curIndex.Cmp(lastBlockHeight) == 0 {
				continue
			}
			if lastBlockHeight.Cmp(big.NewInt(0)) == 0 {
				lastBlockHeight.Set(curIndex)
				continue
			}
			start := big.NewInt(0).Set(lastBlockHeight)
			end := big.NewInt(0).Set(curIndex)
			m.blockListen(start, end)
			lastBlockHeight.Set(curIndex)
			err = m.hScan.SaveHeight(m.ctx, curIndex)
			if err != nil {
				m.logger.WithField(FieldTag, "saveHeight").Error(err)
			}
		case <-m.ctx.Done():
			m.logger.WithField(FieldTag, "close").Info()
			return
		}
	}
}

func (m *monitor) Cancel() {
	m.cancel()
}

func (m *monitor) getBlockHeight() (cur, highest *BlockHeight, err error) {
	sync, err := m.cli.SyncProgress(m.ctx)
	if err != nil {
		return nil, nil, err
	}
	if sync == nil {
		block, err := m.cli.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, nil, err
		}
		return block.Number, block.Number, nil
	} else {
		return big.NewInt(0).SetUint64(sync.CurrentBlock), big.NewInt(0).SetUint64(sync.HighestBlock), nil
	}
}

func (m *monitor) blockListen(start, end *BlockHeight) {
	for i := big.NewInt(0).Set(start); i.Cmp(end) < 0; i.Add(i, big.NewInt(1)) {
		var (
			block *types.Block
			err   error
		)
		for { // 失败阻塞，等待节点修复
			block, err = m.cli.BlockByNumber(context.Background(), i)
			if err != nil {
				m.logger.WithField(FieldTag, "blockByNumber").Errorf("height:%v error:%s", i.String(), err)
				time.Sleep(time.Second)
				continue
			}
			break
		}
		if block.Transactions().Len() > 0 {
			m.analyzeBlock(block)
		}

	}
}

func (m *monitor) analyzeBlock(block *types.Block) {
	for _, v := range block.Transactions() {
		signer := types.NewLondonSigner(v.ChainId())
		msg, err := v.AsMessage(signer, nil)
		if err != nil {
			m.logger.WithField(FieldTag, "asMessage").Error(err)
			continue
		}
		if msg.To() == nil {
			continue
		}
		if m.handler.ContainContact(m.ctx, msg.To().Hex()) {
			txInfo, err := m.analyzeTx(v.Hash(), &msg)
			if err != nil {
				m.logger.WithField(FieldTag, "analyzeTx").Error(err)
				continue
			}
			if txInfo != nil {
				m.handler.Do(m.ctx, txInfo)
			}
		}
	}
}

func (m *monitor) analyzeTx(txHash common.Hash, msg *Message) (*TxInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			m.logger.WithField(FieldTag, "analyzeTx").Errorf("panic cover err:%v", err)
		}
	}()
	txInfo, isPending, err := m.cli.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, err
	}
	if isPending {
		return nil, nil
	}
	txRe, errTxRe := m.cli.TransactionReceipt(context.Background(), txHash)
	fee := big.NewInt(0)
	if errTxRe == nil && txRe != nil {
		fee = fee.SetUint64(txRe.GasUsed)
		fee = fee.Mul(fee, txInfo.GasPrice())
	}
	act, err := m.decoder.DecodeTxData(msg.Data())
	if err != nil {
		return nil, err
	}
	ti := &TxInfo{
		Message: msg,
		Receipt: txRe,
		Action:  act,
		TxHash:  txHash.Hex(),
		Fee:     fee,
		Height:  txRe.BlockNumber,
		Status:  txRe.Status == 1,
	}
	return ti, nil
}
