package main

import (
	"ETL-Ethereum/conf"
	monitor2 "ETL-Ethereum/monitor"
	"time"
)

func main1() {
	opt := &monitor2.Options{
		Config: conf.Config{

		},
		RpcUrl: "https://eth-goerli.alchemyapi.io/v2/2bLTULQFUA8CiZgEws9Mhtt0TRXkBBRf",
	}

	// monitor
	monitor, err := monitor2.New(opt)
	if err != nil {
		panic(err)
	}

	monitor.Run()
	time.Sleep(time.Second * 30)
	monitor.Cancel()
}
