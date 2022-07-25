package main

import (
    "fmt"
    "net/http"
    "os"
	"encoding/json"
    "bytes"
)

const COUNT_ARGS = 3
const TOKEN = "*"
const TG_API_BASE = "https://api.telegram.org/bot" + TOKEN  + "/"

type config struct{
	ip string
	port string
}

func (receiver *config) getFullAddress() string {
	return receiver.ip + ":" + receiver.port
}
 
func main() {
    config := readConfig()

    setHandlers()

    fmt.Println("Server is listening " + config.getFullAddress())
    http.ListenAndServe(config.getFullAddress(), nil)
}

func readConfig() config {
	args := os.Args
	if len(args) < COUNT_ARGS {
		panic("All settings aren't present")
	}
	return config{ ip: os.Args[1], port: os.Args[2] }
}

func setHandlers() {
	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
        var update Update
        decoder := json.NewDecoder(r.Body)
        decoder.DisallowUnknownFields()
        decoder.Decode(&update)

        sendMessageToChat(update.Message.Chat.Id, update.Message.Text)
    })
}

func sendMessageToChat(chat_id int, text string) {
    apiUrl := TG_API_BASE + "sendMessage"
    request := SendMessageRequest{ Text : text, ChatId : chat_id }
    data, _ := json.Marshal(request)
    sendPostJsonRequest(apiUrl, data)
}

func sendPostJsonRequest(url string, json_data []byte) {
    request, _ := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
    request.Header.Set("Content-Type", "application/json; charset=UTF-8")

    client := &http.Client{}
    client.Do(request)
    fmt.Println()
}