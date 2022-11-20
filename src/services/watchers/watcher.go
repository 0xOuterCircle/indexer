package watchers

import (
    "github.com/0xOuterCircle/indexer/src/config"
    "log"
)

func Run() {
    network := config.Global.Network
    log.Println("Watcher: initializing")

    if network == "evm" {
        (&factoryNewDaoWatcherEvm{}).Run()
    } else {
        log.Fatal("wrong global config. No wathcers match provided chain:", network)
    }

    log.Println("Watcher: running")
}