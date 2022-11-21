package notifier

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "os"
)

var chatGeneral = os.Getenv("NOTIFIER_CHAT_GENERAL")

type notifierMessage struct {
    Text                    string `json:"text"`
    ChatId                  string `json:"chat_id"`
    ParseMode               string `json:"parse_mode"`
    DisableNotifications    bool   `json:"disable_notifications"`
}


var client = &http.Client{}
var sendEndpoint = "https://api.telegram.org/bot" + os.Getenv("NOTIFIER_TOKEN") + "/sendMessage"
func send(message notifierMessage) {
    payload, err := json.Marshal(message)
    req, err := http.NewRequest("POST", sendEndpoint, bytes.NewBuffer(payload))
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    err = resp.Body.Close()
    if err != nil { log.Println(err) }
}