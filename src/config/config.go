package config

import (
    "encoding/json"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
    "os"
)

type Config struct {
    EnvironmentId    int
    EnvironmentTitle string
    Network          string         // evm
    ChainID          *big.Int       // 1/5
    FactoryAddress   common.Address
    ApiUrl           string

    // Do not edit JSON tags below!
    DbConnStr        string
    RpcWs            string
    RpcHttp          string
    RpcClient        *ethclient.Client
}

var Global *Config

func Init()  {
    Global = ethDevConfig
}

// Presets

var ethDevConfig = &Config{
    EnvironmentId:    1,
    EnvironmentTitle: "Goerly",
    Network:          "evm",
    ChainID:          big.NewInt(5),
    FactoryAddress:   common.HexToAddress("0x272d3c265600f99f024451d2cdfd48e8e549dcba"),
    ApiUrl:           "https://dolphin-app-mays8.ondigitalocean.app",

    DbConnStr:        os.Getenv("DB_DEV"),
    RpcWs:            os.Getenv("RPC_WS"),
    RpcHttp:          os.Getenv("RPC_HTTP"),
    RpcClient:        nil,
}

func (c *Config) MarshalJSON() (resp []byte, err error) {
    type Alias Config
    resp, err = json.Marshal(&struct {
        EnvironmentId    int               `json:"environment_id"`
        EnvironmentTitle string            `json:"environment_title"`
        Network          string            `json:"network"`         // evm
        ChainID          *big.Int          `json:"chain_id"`        // 1/5
        FactoryAddress   common.Address    `json:"factory_address"`
    }{
        EnvironmentId: c.EnvironmentId,
        EnvironmentTitle: c.EnvironmentTitle,
        Network: c.Network,
        ChainID: c.ChainID,
        FactoryAddress: c.FactoryAddress,
    })
    return
}