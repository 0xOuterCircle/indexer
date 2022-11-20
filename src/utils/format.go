package utils

import (
    "github.com/0xOuterCircle/indexer/src/config"
    "github.com/ethereum/go-ethereum/common"
)

func HashToAddressString(hash common.Hash) string {
    if config.Global.Network == "evm" {
        return common.HexToAddress(hash.Hex()).Hex()
    }
    return ""
}