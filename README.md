# ETL-Ethereum

## init .env

```
export MYSQL_URL        =   
export MYSQL_Scheme     =    
export MYSQL_User       =   
export MYSQL_PASSWORD   =   
export ETH_RPC_URL      = 

source .env
```

## CREATE TABLE

go run cmd/main --init

## SET CONFIG

```
{
    "ChainConfig": {
      "BackwardBlockNumber": 5
    },
        "DBConfig": {
            "Debug": true,
    },
    "ChainListenConfig": {
        "ListenSlot": 5,
        "BatchSize": 10,  // 开启的协程数量
        "defer": 12,      //  区块的确认数
        "From": 0,        //  区块开始同步   from 和 to必须同时定义!
        "To": 0,          //  同步结束的高度  0 < from + BatchSize < to
    }
}

# 默认为增量同步数据
```

## SET TOKEN LIST

```
{
    "tokenList": [
        {
            "name": "Wrapped Ether",
            "symbol": "WETH",
            "address": "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6",
            "decimal": 18
        },
        {
            "name": "Uniswap",
            "symbol": "UNI",
            "address": "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
            "decimal": 18
        }
    ]
}
```

## SET INITIAL PARAMS

```
# example: t_chain
INSERT INTO `t_chain` VALUES(1, 14600000, 5, "Ethereum", 1, '2022-04-24 17:00:01', '2022-04-24 17:00:01');

# eample: t_chain_info
INSERT INTO `t_chain_info` VALUES ('1', 'Ethereum', 'Goerli', '5', 'https://goerli.etherscan.io', 'https://goerli.infura.io/v3/', '0x7066608eF82ff3e7B24eFF45C871712F612345', '0x7066608eF82ff3e7B24eFF45C871712F612345', '0x7066608eF82ff3e7B24eFF45C871712F612345', '0x7066608eF82ff3e7B24eFF45C871712F612345', '2022-03-17 17:00:01', '2022-03-17 17:00:01', 'ETH', 'ws://127.0.0.1:8546', '0x7066608eF82ff3e7B24eFF45C871712F612345', '0x7066608eF82ff3e7B24eFF45C871712F612345', '0x7066608eF82ff3e7B24eFF45C871712F612345', '0.0100');
```

