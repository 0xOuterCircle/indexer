package watchers

import (
    "context"
    "github.com/0xOuterCircle/indexer/src/config"
    "github.com/0xOuterCircle/indexer/src/entities/dao"
    "github.com/0xOuterCircle/indexer/src/utils"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "log"
)

type factoryNewDaoWatcherEvm struct {}

func (fw *factoryNewDaoWatcherEvm) Run() {

    log.Println("FactoryNewDaoWatcherEVM: dialing node")
    client, err := ethclient.Dial(config.Global.RpcWs)
    if err != nil { log.Fatal(err) }
    log.Println("FactoryNewDaoWatcherEVM: connected to node")

    query := ethereum.FilterQuery{
        Addresses: []common.Address{config.Global.FactoryAddress},
        Topics: [][]common.Hash {{
            common.HexToHash("0xdcdaa9087962b773cd00b3088349710fd7f4c555341f3d594ac9a7a8f1a5e9ea"),
        }},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }
    go fw.worker(sub, logs)
}

func (fw *factoryNewDaoWatcherEvm) worker(sub ethereum.Subscription, logs <-chan types.Log) {
    for {
        select {
            case err := <-sub.Err():
                log.Println(err)  // TODO Sentry
            case logEntry := <-logs:

                newDao := dao.Dao{
                    Id: nil,
                    Parent: nil,
                    Registry: utils.HashToAddressString(logEntry.Topics[1]),
                    Governance: utils.HashToAddressString(logEntry.Topics[2]),
                    Creator: "",
                }

                err := newDao.Create()
                if err != nil {
                    log.Println(err)
                }
        }
    }
}