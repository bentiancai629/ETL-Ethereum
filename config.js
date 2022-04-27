require("dotenv").config();
const stripHexPrefix = require('strip-hex-prefix');
const Web3 = require('web3');
const crowdSwapJson = require('./build/contracts/CrowdSwap.json');
const uniswapPriceOracleJson = require('./build/contracts/UniswapPriceOracle.json');
const chainlinkPriceOracleJson = require('./build/contracts/ChainlinkPriceOracle.json');
const customPriceOracleJson = require('./build/contracts/CustomPriceOracle.json');
const chainlinkAggregatorV3ContractJson = require('./build/contracts/AggregatorV3Interface.json');
const CryptoJS = require("crypto-js");
const tripledes = require("crypto-js/tripledes");
const decrypt = (cipherText, password) => tripledes.decrypt(cipherText, password).toString(CryptoJS.enc.Utf8);
const encrypt = (plainText, password) => tripledes.encrypt(plainText, password).toString();

// giraffe daring february life stay oppose level letter melody coin vague sick

// 0xCc9caE3d07A7De1A7a60587B22A1EF607E4c4AFE
// 0x96994fdb018d181003909763b1e099e145b07164401ee755ed7d2e6391840adc

let accounts = {
    contractOwner: {
        // address: "0xfEC499fb214592E79BAD92d6E5bf2E084F69fDe5", // 0x318d754fa5a6ef1cd0a524b6c488df288e7fe34a13d1fe7328cfc3284cc13484
        // privateKey: "U2FsdGVkX18306I6hGwFZ50FzOW2y3RxlbjWfz/uM5tB4Q9PJMxVjjDeZCnRExddwBo5k4iURJxYtlEGYGgM4KCfEIrAeHRpYKVoFNOIaBipo96FPb/IMQ==",
        // password: "1qaz@WSX",

        // qeth
        // address: "0xfEC499fb214592E79BAD92d6E5bf2E084F69fDe5", // 0x318d754fa5a6ef1cd0a524b6c488df288e7fe34a13d1fe7328cfc3284cc13484
        // privateKey: "0x318d754fa5a6ef1cd0a524b6c488df288e7fe34a13d1fe7328cfc3284cc13484"

        // epot
        address: "0x17EAD008c360c4730ADedf8C4780E84F412Baa28",  // EPOT
        privateKey: "dbd1635045180e86fb3e4d4dc44eb28ba0ded6daf0f2bdb55c9ceb8aefe90393",
    },
    etherStaker1: {
        address: "0x8078BF1B18f56DFc5a86d73dEf1f3B32957bCf7d",
        privateKey: "0xca9d26b45389d3b1f0ae5abca31fbda5b59fb0db05eeba4e13ec15a1600441d7",
    },
    etherStaker2: {
        address: "0x1059F2325d47Ef64529F9c61cc222E2a1a969646",
        privateKey: "0x2ce596ceedbaf484079392ae634534203a6c429b3331a4ad8e116d36ec7a61ca",
    },
    feeAccount: {
        address: "0x19DF057056847A5016abb3c127f7907bec266250",
        privateKey: "b45dc2d50f6179991ecbc3f0337038a1d4fecdfa2595dc4f2594d3979122ff83",
    },
    etherWallet: {
        address: process.env.ETH_HOT_WALLET_ADDRESS,
        privateKey: "8910fbaaedba69ce34d16e4b507aceff6894e8f60813a6cc739b6766c81e5df3"
    },
    proxyAdmin: {
        address: "0x1059F2325d47Ef64529F9c61cc222E2a1a969646",
        privateKey: "0x2ce596ceedbaf484079392ae634534203a6c429b3331a4ad8e116d36ec7a61ca"
    },
};

let web3 = new Web3(process.env.ETH1_NODE_RPC);

let config = {
    crowdSwapContract: {
        deployArgs: {
            swapPrice: web3.utils.toWei("1", "ether"),
            wallet: accounts.etherWallet.address,
            usdFeePerETH: web3.utils.toWei("1", "ether"),
            feeWallet: accounts.feeAccount.address,
            oracle: process.env.CUSTOM_PRICE_ORACLE_CONTRACT_ADDRESS,
            tokenName: "EPotter Ether Staking Liquid Token",
            tokenSymbol: "EPOT",
        },
        abi: crowdSwapJson.abi,
        bytecode: stripHexPrefix(crowdSwapJson.bytecode),
        address: process.env.SWAP_CONTRACT_ADDRESS
    },
    uniswapPriceOracleContract: {
        deployArgs: {
            usdAddress: process.env.UNISWAP_PRICE_ORACLE_USD_ADDRESS,
        },
        abi: uniswapPriceOracleJson.abi,
        bytecode: stripHexPrefix(uniswapPriceOracleJson.bytecode),
        address: process.env.UNISWAP_PRICE_ORACLE_CONTRACT_ADDRESS
    },
    chainlinkPriceOracleContract: {
        deployArgs: {
            aggregatorAddress: process.env.CHAINLINK_PRICE_ORACLE_AGGREGATOR_ADDRESS,
            effectiveMinutes: 120
        },
        abi: chainlinkPriceOracleJson.abi,
        bytecode: stripHexPrefix(chainlinkPriceOracleJson.bytecode),
        address: process.env.CHAINLINK_PRICE_ORACLE_CONTRACT_ADDRESS
    },
    customPriceOracleContract: {
        deployArgs: {
            ethPrice: web3.utils.toWei("2500.00", "ether")
        },
        abi: customPriceOracleJson.abi,
        bytecode: stripHexPrefix(customPriceOracleJson.bytecode),
        address: process.env.CUSTOM_PRICE_ORACLE_CONTRACT_ADDRESS
    },
    chainlinkAggregatorV3Contract: {
        abi: chainlinkAggregatorV3ContractJson.abi,
        bytecode: stripHexPrefix(chainlinkAggregatorV3ContractJson.bytecode),
        address: process.env.CHAINLINK_PRICE_ORACLE_AGGREGATOR_ADDRESS
    },
    depositContract: {
        abi: [
            {"inputs": [], "stateMutability": "nonpayable", "type": "constructor"}, {
                "anonymous": false,
                "inputs": [{
                    "indexed": false,
                    "internalType": "bytes",
                    "name": "pubkey",
                    "type": "bytes"
                }, {
                    "indexed": false,
                    "internalType": "bytes",
                    "name": "withdrawal_credentials",
                    "type": "bytes"
                }, {"indexed": false, "internalType": "bytes", "name": "amount", "type": "bytes"}, {
                    "indexed": false,
                    "internalType": "bytes",
                    "name": "signature",
                    "type": "bytes"
                }, {"indexed": false, "internalType": "bytes", "name": "index", "type": "bytes"}],
                "name": "DepositEvent",
                "type": "event"
            }, {
                "inputs": [{"internalType": "bytes", "name": "pubkey", "type": "bytes"}, {
                    "internalType": "bytes",
                    "name": "withdrawal_credentials",
                    "type": "bytes"
                }, {"internalType": "bytes", "name": "signature", "type": "bytes"}, {
                    "internalType": "bytes32",
                    "name": "deposit_data_root",
                    "type": "bytes32"
                }], "name": "deposit", "outputs": [], "stateMutability": "payable", "type": "function"
            }, {
                "inputs": [],
                "name": "get_deposit_count",
                "outputs": [{"internalType": "bytes", "name": "", "type": "bytes"}],
                "stateMutability": "view",
                "type": "function"
            }, {
                "inputs": [],
                "name": "get_deposit_root",
                "outputs": [{"internalType": "bytes32", "name": "", "type": "bytes32"}],
                "stateMutability": "view",
                "type": "function"
            }, {
                "inputs": [{"internalType": "bytes4", "name": "interfaceId", "type": "bytes4"}],
                "name": "supportsInterface",
                "outputs": [{"internalType": "bool", "name": "", "type": "bool"}],
                "stateMutability": "pure",
                "type": "function"
            }],
        address: process.env.DEPOSIT_CONTRACT_ADDRESS,
    },
    accounts: accounts,
    web3: web3,
    decrypt: decrypt,
    encrypt: encrypt
};

(async () => {
    await Promise.all(Object.keys(accounts).map(async (k) => {
        if (accounts[k].privateKey !== undefined) {
            let pri = accounts[k].privateKey;
            if (accounts[k].password !== undefined && accounts[k].password !== "") {
                pri = decrypt(pri, accounts[k].password);
            }
        accounts[k].account = web3.eth.accounts.wallet.add(pri);
        }
    }));
})();

module.exports = config;
