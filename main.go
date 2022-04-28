package main

import (
	"ETL-Ethereum/conf"
	"ETL-Ethereum/monitor"
	"time"
)

func main() {
	opt := &monitor.Options{
		Config: conf.Config{

		},
		RpcUrl: "https://eth-goerli.alchemyapi.io/v2/2bLTULQFUA8CiZgEws9Mhtt0TRXkBBRf",
	}

	// monitor
	monitor, err := monitor.New(opt)
	if err != nil {
		panic(err)
	}

	monitor.Run()
	time.Sleep(time.Second * 30)
	monitor.Cancel()
}
