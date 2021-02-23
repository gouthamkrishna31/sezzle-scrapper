package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please provide the correct number of arguments expected 3 found ", len(os.Args)-1)

	}
	crawlLink := string(os.Args[1])
	itemCount, _ := strconv.Atoi(os.Args[2])
	crawl(crawlLink, itemCount)
}

func crawl(crawlLink string, itemCount int) {
	counter := 0
	var masterData []Data
	wg := new(sync.WaitGroup)
	c := colly.NewCollector(colly.AllowedDomains("imdb.com", "www.imdb.com"))
	infoCollector := colly.NewCollector(colly.AllowedDomains("imdb.com", "www.imdb.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Profile URL: ", r.URL.String())
	})
	infoCollector.OnHTML(`script[type="application/ld+json"]`, func(e *colly.HTMLElement) {
		defer wg.Done()
		var imdbdtl ImdbDetails
		err := json.Unmarshal([]byte(e.Text), &imdbdtl)
		if err != nil {
			log.Fatal(err)
		}
		var singleData Data
		singleData.Title = imdbdtl.Name
		singleData.MovieReleaseYear, _ = strconv.Atoi(strings.Split(imdbdtl.DatePublished, "-")[0])
		singleData.ImdbRating, _ = strconv.ParseFloat(imdbdtl.AggregateRating.RatingValue, 1)
		singleData.Summary = imdbdtl.Description
		singleData.Duration = imdbdtl.Duration
		singleData.Genre = imdbdtl.Genre
		masterData = append(masterData, singleData)

	})
	c.OnHTML("td[class=titleColumn]", func(e *colly.HTMLElement) {

		goquerySelection := e.DOM
		link, _ := goquerySelection.Find("a").Attr("href")
		if counter == itemCount {
			return
		}
		wg.Add(1)
		time.Sleep(50 * time.Millisecond)
		go infoCollector.Visit("https://www.imdb.com" + link)
		counter++

	})

	c.Visit(crawlLink)
	wg.Wait()
	output, _ := json.Marshal(masterData)
	fmt.Println(string(output))
}
