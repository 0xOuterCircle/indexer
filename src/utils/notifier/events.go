package notifier

import (
    "github.com/0xOuterCircle/indexer/src/config"
    "os"
)

func Deployment() {
    if os.Getenv("DEV") != "" { return }
    message := notifierMessage{`🚀 `+config.Global.EnvironmentTitle+` Indexer went live!
👉 `+config.Global.ApiUrl, chatGeneral, "markdown", false}
    send(message)
}