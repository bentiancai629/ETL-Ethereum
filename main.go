package main

import (
	"ETL-Ethereum/monitor"
	"time"
)

func main() {
	opt := &monitor.Options{
		RpcUrl: "https://eth-goerli.alchemyapi.io/v2/2bLTULQFUA8CiZgEws9Mhtt0TRXkBBRf",
		AbiStr: "",
		//Handler: handler,
	}

	// monitor
	monitor, err := monitor.New(opt)
	if err != nil {
		panic(err)
	}

	monitor.Run()
	time.Sleep(time.Second * 10)

}
