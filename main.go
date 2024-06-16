package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
)

type Options struct {
	token  string
	domain string
	debug  bool
}

func main() {
	option := loadOptions()
	purge(option)
}

func purge(option Options) {
	log.Println("[purge] Purging cache...")

	if option.debug {
		log.Println("[purge] With options: ", option)
	}

	body := []byte(`{
		"purge_everything": true
	}`)

	url := "https://api.cloudflare.com/client/v4/zones/" + option.domain + "/purge_cache"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+option.token)

	client := &http.Client{}
	res, err := client.Do(req)

	if option.debug == true {
		log.Println("Response: ", res)
		log.Println("Response: ", err)
	}

	if err != nil {
		log.Println("[purge] error during request")
	}

	if res.StatusCode == http.StatusBadRequest {
		log.Println("[purge] Validate if correct token was provided")
	}

	if res.StatusCode == http.StatusNotFound {
		log.Println("[purge] Validate if correct domain was provided")
	}

	if res.StatusCode == http.StatusOK {
		log.Println("[purge] Cache cleaned")
	}

}

func loadOptions() Options {
	log.Println("[loadOptions] Loading options...")
	log.Println("[loadOptions] TOKEN:", os.Getenv("TOKEN"))
	log.Println("[loadOptions] DOMAIN:", os.Getenv("DEBUG"))

	debug := false
	if os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG") == "TRUE" {
		debug = true
	}

	option := Options{token: os.Getenv("TOKEN"), domain: os.Getenv("DOMAIN"), debug: debug}
	return option
}
