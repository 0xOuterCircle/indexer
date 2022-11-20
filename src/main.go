package main

import (
    "github.com/0xOuterCircle/indexer/src/config"
    "github.com/0xOuterCircle/indexer/src/services/api"
    "github.com/0xOuterCircle/indexer/src/services/db"
    "github.com/0xOuterCircle/indexer/src/services/watchers"
)

func main() {

    // Preparing Configuration
    config.Init()

    // Creating DB Instance
    db.Instance().Connect()

    // Initializing Watchers
    watchers.Run()

    // Running API service
    api.Run()
}