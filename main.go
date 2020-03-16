package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
	"github.com/tipej/CoronaApp-Api/models"
	"github.com/tipej/CoronaApp-Api/resources"
)

var newslist = []models.News{}
var fp = gofeed.NewParser()

func newsLink(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.NewEncoder(w).Encode(newslist)
}

// Loop The News Requests
func requestNewsLoop() {
	ticker := time.NewTicker(30 * time.Second)
	for t := range ticker.C {
		_ = t
		go requestNews()
	}
}

// Request News Objects from all of the RSS Feeds
func requestNews() {
	newslist = []models.News{}
	fmt.Println("Fetching News Objects...")
	for k, v := range resources.NewsSources {
		response, err := http.Get(v.RssUrl)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			body, err := ioutil.ReadAll(response.Body)
			_ = err
			feed, err := fp.ParseString(string(body))
			for _, item := range feed.Items {
				equals := false
				if v.Keyword != "" {
					for _, category := range item.Categories {
						if category == v.Keyword {
							equals = true
						}
					}
				} else {
					equals = true
				}
				if equals {
					published, _ := time.Parse(time.RFC1123Z, item.Published)
					newslist = append(newslist, models.News{
						Title:       item.Title,
						Url:         item.Link,
						Description: item.Description,
						Source:      k,
						PubDate:     published.UTC(),
					})
				}
			}
		}
	}
	// Sort news objects based on age
	sort.Slice(newslist, func(i, j int) bool {
		return newslist[i].PubDate.After(newslist[j].PubDate)
	})
}

func main() {
	go requestNewsLoop()
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/news", newsLink)
	log.Fatal(http.ListenAndServe(port, router))
}
