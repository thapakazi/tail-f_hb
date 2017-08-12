package main

import (
	"fmt"
	"github.com/namsral/flag"
	"io/ioutil"
	"net/http"
)

type Fault struct {
	Id, ProjectId int
	Message       string
	Url           string
	Env           string
}

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

func construct_url()(url string) {
	url = api_url + "/projects/" + project_id + "/faults?auth_token=" + auth_token + "&q=environment%3A" + env + "&order=" + order + "&limit=" + limit
	// fmt.Println(url)
	return
}
func main() {
	
	url := construct_url()

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
	// fmt.Println("get:\n", keepLines(string(body), 3))
	fmt.Println(string(body))

}
