package main

import (
	"fmt"
	"github.com/cskr/pubsub"
)

// 区块链处理
type BlockEvent interface {}

// 区块信息
type BlockInfo struct {
	BlockNumber    int64
	BlockTxCounter int64
}

// tx转账
type Transformer interface {
	Transforms(event BlockEvent)
}

// 合约事件
type ContractEventEvent struct {}
// 拿到event
func (c *ContractEventEvent) GetEvent() (cc *ContractEventEvent) {
	d := new(ContractEventEvent)
	return d
}
// 合约-转账事件
type ContractEventTransformer struct {}
func (c *ContractEventTransformer) Transform(event ContractEventEvent) (cEvent *ContractEventEvent) {
	tx := new(ContractEventEvent)
	return tx
}

// TxTransformer
type TxTransformer struct{}
type TxEvent struct{}
func (t *TxTransformer) Transform(event TxEvent) (txEvent *TxEvent) {
	tx := new(TxEvent)
	return tx
}


//  解析器句柄
type Parser interface {
	Parse(event ContractEventEvent)
}
// erc20解析器
type ERC20Parser struct{}
// erc721解析器
type ERC721Parser struct{}
//  erc20解析实现
func (erc20 *ERC20Parser) Parse(event ContractEventEvent) {
	//do save db
}
func (erc721 *ERC721Parser) Parse(event ContractEventEvent) {
	//do save db
}


// erc20/erc721的Transform接口实现
type ERC20TokenTransformer struct {}
func (erc20 *ERC20TokenTransformer) Transform(event ContractEventEvent) (p *Parser) {
	//save database object
	return nil
}
type ERC721TokenTransformer struct {}
func (erc721 *ERC721TokenTransformer) Transform(event ContractEventEvent) (p *Parser) {
	return nil
}

// emit事件
func emitEvent() {}

const topic = "BlockEvent"

func main() {
	// 0. 初始化pubsub
	ps := pubsub.New(0)
	ch := ps.Sub(topic)
	// 1. sub BlockEvENT
	blockEvt := &BlockInfo{100, 10}

	// 2. sub BlockEvENT
	go publishBlockEvt(ps, blockEvt)

	for i := 1; ; i++ {
		if i == 5 {
			go ps.Unsub(ch, "topic")
		}

		if msg, ok := <-ch; ok {
			fmt.Printf("Received %s, %d times.\n", msg, i)
		} else {
			break
		}
	}

	defer ps.Close(topic)
}

func publishBlockEvt(ps *pubsub.PubSub, blockEvt *BlockInfo) {
	for {
		ps.Pub(blockEvt.BlockNumber, topic)
		fmt.Println("blockEvt :", blockEvt.BlockNumber)
	}
}
