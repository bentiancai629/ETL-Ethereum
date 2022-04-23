package listener

import (
	"ETL-Ethereum/conf"
	"ETL-Ethereum/contracts/erc20"
	"ETL-Ethereum/contracts/erc721"
	"ETL-Ethereum/contracts/exchange"
	"ETL-Ethereum/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	//"github.com/shopspring/decimal"

)

type ChainListenCore struct {
	chainInfo models.ChainInfo
	gethCfg   *conf.ChainListenConfig
	rpcClient *rpc.Client
	rawClient *ethclient.Client
	mu        sync.Mutex
}

func NewChainListenCore(chainInfo models.ChainInfo, gethCfg *conf.ChainListenConfig) (*ChainListenCore, error) {
	rpcClient, err := rpc.Dial(chainInfo.InternalRpcURL)
	if rpcClient == nil || err != nil {
		return nil, fmt.Errorf("rpc client works error(%s)", err)
	}
	rawClient, err := ethclient.Dial(chainInfo.InternalRpcURL)
	if rawClient == nil || err != nil {
		return nil, fmt.Errorf("raw client works error(%s)", err)
	}
	return &ChainListenCore{
		gethCfg:   gethCfg,
		chainInfo: chainInfo,
		rpcClient: rpcClient,
		rawClient: rawClient,
		mu:        sync.Mutex{},
	}, nil
}

func (c *ChainListenCore) GetChainListenSlot() uint64 {
	return c.gethCfg.ListenSlot
}

func (c *ChainListenCore) GetBatchSize() uint64 {
	return c.gethCfg.BatchSize
}

func (c *ChainListenCore) GetDefer() uint64 {
	return c.gethCfg.Defer
}

func (c *ChainListenCore) GetCurrentBlockHeight() (uint64, error) {
	var result hexutil.Big
	err := c.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
}

func (c *ChainListenCore) GetLatHeaderByNumber(number uint64) (*PlatonHeader, error) {
	var header *PlatonHeader
	err := c.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(number))), true)
	return header, err
}

func (c *ChainListenCore) GetEthHeaderByNumber(number uint64) (*types.Header, error) {
	header := &types.Header{}
	err := c.rpcClient.CallContext(context.Background(), header, "eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(number))), false)
	return header, err
}

// core
func (c *ChainListenCore) HandleNewBlock(height uint64) ([]*models.Erc20TransferEvent, error) {
	var erc20Evts []*models.Erc20TransferEvent

	header, err := c.GetEthHeaderByNumber(height)
	if err != nil {
		return nil, err
	}
	if header == nil {
		return nil, fmt.Errorf("there is no geth block")
	}


	//todo
	height = 6761665

	erc20Evt, err := c.getERC20EventByBlockNumber(c.gethCfg.WETHAddr, height,height, 2) // todo enum
	if err != nil {
		return nil, err
	}
	erc20Evts = append(erc20Evts, erc20Evt...)

	for _, item := range erc20Evts {
		logs.Info("(claim) from chain, txhash: %s", item.Hash)
	}

	return erc20Evts, nil
}

// core
func (c *ChainListenCore) HandleNewBlockBk(height uint64) ([]*models.Erc20TransferEvent, error) {
	var erc20Evts []*models.Erc20TransferEvent

	header, err := c.GetEthHeaderByNumber(height)
	if err != nil {
		return nil, err
	}

	if header == nil {
		return nil, fmt.Errorf("there is no geth block")
	}

	erc20Evt, err := c.getERC20EventByBlockNumber(c.gethCfg.WETHAddr, height, height, 1)
	if err != nil {
		return nil, err
	}
	erc20Evts = append(erc20Evts, erc20Evt...)

	for _, item := range erc20Evts {
		logs.Info("(claim) from chain, txhash: %s", item.Hash)
	}

	return erc20Evts, nil
}

// contractCall 处理 lido的合约地址
func (c *ChainListenCore) get721EventByBlockNumber(erc721Addr string, startHeight, endHeight uint64, eventType uint8) ([]*models.Erc20TransferEvent, error) {
	nftAddr := common.HexToAddress(erc721Addr)
	erc721Instance, err := erc721.NewErc721(nftAddr, c.rawClient)

	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := getFilterOpts(startHeight, endHeight)
	// get geth transfer events from given block
	claimTransactions := make([]*models.Erc20TransferEvent, 0)
	//fromFilter := []common.Address{common.HexToAddress("0x0000000000000000000000000000000000000000")}
	transferEvent, err := erc721Instance.FilterTransfer(opt, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, filter transfer events :%s", err.Error())
	}
	for transferEvent.Next() {
		evt := transferEvent.Event
		text, err := json.Marshal(evt)
		if err != nil {
			return nil, err
		}
		// 重新封装了对象
		claimTx := getNFtEvent(evt.Raw, eventType, c.chainInfo.ChainName, text)
		claimTransactions = append(claimTransactions, claimTx)
	}

	return claimTransactions, nil
}

// contractCall 处理 lido的合约地址
func (c *ChainListenCore) getERC20EventByBlockNumber(erc20Addr string, startHeight, endHeight uint64, eventType uint8) ([]*models.Erc20TransferEvent, error) {
	erc20AddrHex := common.HexToAddress(erc20Addr)
	erc20Instance, err := erc20.NewErc20(erc20AddrHex, c.rawClient)

	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}

	// from - to height自定义
	opt := getFilterOpts(startHeight, endHeight)
	//get geth transfer events from given block
	erc20Transactions := make([]*models.Erc20TransferEvent, 0)
	transferEvent, err := erc20Instance.FilterTransfer(opt, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, filter transfer events :%s", err.Error())
	}

	for transferEvent.Next() {
		evt := transferEvent.Event
		text, err := json.Marshal(evt)
		if err != nil {
			return nil, err
		}
		// 自定义
		//transfer := getERC20Event(evt.Raw, eventType, c.chainInfo.ChainName, text)

		fmt.Println("txHash ",evt.Raw.TxHash.Hex())

		fmt.Println("transfer  from ", evt.Src)
		fmt.Println("transfer  to ",evt.Dst)
		fmt.Println("transfer  amount ",evt.Wad.Uint64())

		//fAmount := new(big.Float)
		//fAmount.SetString(evt.Wad.String())
		//ethValue := new(big.Float).Quo(fAmount, big.NewFloat(math.Pow10(18)))
		//fmt.Println("transfer  amountInEth ",ethValue)

		transferDao := &models.Erc20TransferEvent{
			Height:    evt.Raw.BlockNumber,
			Hash:      evt.Raw.TxHash.Hex(),
			RawIndex: uint8(evt.Raw.Index),
			EventType: eventType,
			ChainName: c.chainInfo.ChainName,
			From: evt.Src.String(),
			To: evt.Dst.String(),
			Amount: evt.Wad.Uint64(), // todo 需要处理decimal
			EventJson: text,
			Status:    0,
		}

		erc20Transactions = append(erc20Transactions, transferDao)
	}

	return erc20Transactions, nil
}

func getFilterOpts(startHeight, endHeight uint64) *bind.FilterOpts {
	return &bind.FilterOpts{
		Start:   startHeight,
		End:     &endHeight,
		Context: context.Background(),
	}
}

//  处理evt
func getNFtEvent(raw types.Log, eventType uint8, chainName string, text []byte) *models.Erc20TransferEvent {
	return &models.Erc20TransferEvent{
		Height:    raw.BlockNumber,
		Hash:      raw.TxHash.Hex(),
		EventType: eventType,
		ChainName: chainName,
		EventJson: text,
		Status:    0,
	}
}

//  处理evt
func getERC20Event(raw types.Log, eventType uint8, chainName string, text []byte) *models.Erc20TransferEvent {
	return &models.Erc20TransferEvent{
		Height:    raw.BlockNumber,
		Hash:      raw.TxHash.Hex(),
		EventType: eventType,
		ChainName: chainName,
		//From: raw.Address,
		//To: raw.Address,
		EventJson: text,
		Status:    0,
	}
}

func (c *ChainListenCore) getExchangeEventByBlockNumber(erc721Addr string, startHeight, endHeight uint64) ([]*models.Erc20TransferEvent, error) {
	nftAddr := common.HexToAddress(erc721Addr)
	erc721Instance, err := exchange.NewExchange(nftAddr, c.rawClient)

	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}
	opt := getFilterOpts(startHeight, endHeight)
	// get geth transfer events from given block
	claimTransactions := make([]*models.Erc20TransferEvent, 0)
	//fromFilter := []common.Address{common.HexToAddress("0x0000000000000000000000000000000000000000")}
	matchedEvent, err := erc721Instance.FilterOrdersMatched(opt, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, filter transfer events :%s", err.Error())
	}
	for matchedEvent.Next() {
		evt := matchedEvent.Event
		text, err := json.Marshal(evt)
		if err != nil {
			return nil, err
		}
		claimTx := getNFtEvent(evt.Raw, 1, c.chainInfo.ChainName, text)
		claimTransactions = append(claimTransactions, claimTx)
	}

	cancelledEvent, err := erc721Instance.FilterOrderCancelled(opt, nil)
	if err != nil {
		return nil, fmt.Errorf("GetSmartContractEventByBlock, filter transfer events :%s", err.Error())
	}
	for cancelledEvent.Next() {
		evt := cancelledEvent.Event
		text, err := json.Marshal(evt)
		if err != nil {
			return nil, err
		}
		claimTx := getNFtEvent(evt.Raw, 3, c.chainInfo.ChainName, text)
		claimTransactions = append(claimTransactions, claimTx)
	}
	return claimTransactions, nil
}


// ToDecimal wei to decimals
//func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
//	value := new(big.Int)
//	switch v := ivalue.(type) {
//	case string:
//		value.SetString(v, 10)
//	case *big.Int:
//		value = v
//	}
//
//	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
//	num, _ := decimal.NewFromString(value.String())
//	result := num.Div(mul)
//
//	return result
//}