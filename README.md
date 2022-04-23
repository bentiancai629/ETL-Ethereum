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
        "URL": "",
        "Scheme": "",
        "User": ""
    },
    "ChainListenConfig": {
        "ListenSlot": 5,
        "BatchSize": 1,  //  开启的协程
        "defer": 12,     //  考虑链上分叉的区块影响
        "From": 0,       //  区块开始同步   from 和 to必须同时定义!
        "To": 0,         //  同步结束的高度  0 < from + BatchSize < to
        "Url": "wss://goerli.infura.io/ws/v3/ba851582c0314accaebcde9010e50e83"
    }
}
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

