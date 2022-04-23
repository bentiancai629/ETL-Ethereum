package listener

import (
	"math"
	"runtime/debug"
	"time"

	"github.com/beego/beego/v2/core/logs"

	"ETL-Ethereum/conf"
	"ETL-Ethereum/dao"
	"ETL-Ethereum/models"
)

type ChainListen struct {
	core   *ChainListenCore
	db     *dao.DataBase
	height uint64
	exit   chan bool
}

var chainListenAry []*ChainListen

func StartLandListen(dbCfg conf.DBConfig, chain *conf.ChainListenConfig) {
	daoLand := dao.NewLandDao(dbCfg)
	if daoLand == nil {
		panic("sql server is valid")
	}
	list := daoLand.GetChainInfo()
	for _, chainInfo := range list {
		go ListenCore(daoLand, chainInfo, chain)
	}
}

// 链上监听
func ListenCore(dao *dao.DataBase, chainInfo models.ChainInfo, chain *conf.ChainListenConfig) {
	logs.Info("ListenCore :%v", chainInfo)
	core, err := NewChainListenCore(chainInfo, chain)
	if err != nil {
		panic(err)
	}

	chainListen := &ChainListen{
		core: core,
		db:   dao,
		exit: make(chan bool, 0),
	}

	chainListenAry = append(chainListenAry, chainListen)
	for {
		exit := chainListen.listenChain()
		if exit {
			close(chainListen.exit)
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func StopLandListen() {
	for _, ch := range chainListenAry {
		ch.exit <- true
	}

}


// handel height
func (cl *ChainListen) HandleNewBlock(height uint64) (ct []*models.Erc20TransferEvent, err error) {
	ct, err = cl.core.HandleNewBlock(height)
	if err != nil {
		logs.Error("Possible inconsistent chain height %d", height)
		return
	}
	return
}


// core logic
func (cl *ChainListen) listenChain() (exit bool) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("service start, recover info: %s", string(debug.Stack()))
			exit = false
		}
	}()
	currentHeight, err := cl.core.GetCurrentBlockHeight() // 当前高度
	if err != nil || currentHeight == 0 {
		panic(err)
	}
	var chainName = cl.core.chainInfo.ChainName

	chain, err := cl.db.GetChain(chainName)  // 数据库高度
	if err != nil {
		panic(err)
	}
	if chain.Height == 0 {
		chain.Height = currentHeight
	}
	cl.db.UpdateChain(chain)
	if cl.height != 0 {
		chain.Height = cl.height
	}
	ticker := time.NewTicker(time.Second * time.Duration(cl.core.GetChainListenSlot()))


	for {
		select {
		case <-ticker.C:
			height, err := cl.core.GetCurrentBlockHeight()
			if err != nil || height == 0 || height == math.MaxUint64 {
				logs.Error("listenChain - cannot get chain height, err: %s", err)
				continue
			}
			if chain.Height >= height-cl.core.GetDefer() {
				continue
			}
			logs.Info("ListenChain - chain latest height is %d, listen height: %d", height, chain.Height)
			for chain.Height < height-cl.core.GetDefer() {
				batchSize := cl.core.GetBatchSize()
				if batchSize == 0 {
					batchSize = 1
				}
				if batchSize > height-chain.Height-cl.core.GetDefer() {
					batchSize = height - chain.Height - cl.core.GetDefer()
				}
				ch := make(chan bool, batchSize)

				// 5个 协程
				for i := uint64(1); i <= batchSize; i++ {
					//  并发处理
					go func(height uint64) {
						transactionsDao, err := cl.HandleNewBlock(height)
						if err != nil {
							logs.Error("HandleNewBlock %d err: %v", height, err)
							ch <- false
							return
						}

						// 事件存入数据库
						err = cl.db.SaveEventsIgnoredByHash(transactionsDao)
						if err != nil {
							logs.Error("UpdateEvents on block %d err: %v", height, err)
							ch <- false
						} else {
							ch <- true
						}

					}(chain.Height + i)
				}


				// 必须任务完成
				allTaskSuccess := true

				for j := 0; j < int(batchSize); j++ {
					ok := <-ch
					if !ok {
						allTaskSuccess = false
					}
				}
				close(ch)
				if !allTaskSuccess {
					// 可以加回滚逻辑
					break
				}

				chain.Height += batchSize
				if err := cl.db.UpdateChain(chain); err != nil {
					logs.Error("UpdateChain [height:%d] err %v", chain.Height, err)
					chain.Height -= batchSize
				}
			}
		case <-cl.exit:
			logs.Info("cross chain listen exit")
			return true
		}
	}
}
