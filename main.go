package main

import (
	"encoding/json"
	"github.com/namsral/flag"
	"io/ioutil"
	"net/http"
)

var api_url, env, auth_token, order, limit, project_id string

func init() {
	// fmt.Println("from init(): I am init, I feel blessed")
	flag.StringVar(&api_url, "api_url", "https://app.honeybadger.io/v2", "API URL")
	flag.StringVar(&auth_token, "auth_token", "your_token_from_honeybadger", "Authentication Token from Honeybadger, read is enough")
	flag.StringVar(&env, "env", "production", "Error specific to ENVIRONMENT to scrape")
	flag.StringVar(&project_id, "project_id", "12345", "My Project Id in Honeybadger")
	flag.StringVar(&limit, "limit", "3", "Number of errors to scrape")
	flag.StringVar(&order, "order", "recent", "Order of scraping, available opts[recent,frequent]")
	flag.Parse()

}


func main() {

	url := Construct_Url() // form url, check helpers function

	// get the config
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	faults := Faults{}
	// decode json body first
	json.Unmarshal([]byte(body), &faults)

	PrintFaults(&faults) // print with colors, check print_out functions
}


