package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("search-key.json")
	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/cse")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)
	cseService, err := customsearch.New(client)
	search := cseService.Cse.List("スカサハ")

	// 検索エンジンIDを適宜設定
	search.Cx("xxxxx")
	// Custom Search Engineで「画像検索」をオンにする
	search.SearchType("image")

	search.Start(1)
	call, err := search.Do()
	if err != nil {
		log.Fatal(err)
	}

	for index, r := range call.Items {
		fmt.Println(index)
		fmt.Println(r.Link)
	}
}
