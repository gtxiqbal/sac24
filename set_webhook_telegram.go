package main

import (
	"encoding/json"
	"fmt"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web/telegram/request"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	botTokenTelegram := os.Getenv("BOT_TOKEN_TELEGRAM")
	url := os.Args[1]
	if strings.HasSuffix(url, "/") {
		url = url[0 : len(url)-1]
	}

	get, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s/api/telegram&drop_pending_updates=true",
		botTokenTelegram, url))
	helper.PanicIfError(err)
	defer get.Body.Close()

	response, err := ioutil.ReadAll(get.Body)
	var resp request.ResultResponse
	_ = json.Unmarshal(response, &resp)
	if resp.Ok {
		fmt.Println("Webhook berhasil disetting")
	} else {
		fmt.Println("Webhook gagal disetting: " + resp.Description)
	}
}
