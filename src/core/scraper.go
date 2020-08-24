package core

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gurban/src/dict"
	"io"
	"log"
	"net/http"
	"github.com/fatih/color"
)


func getResponseForURL(url string) (resp io.ReadCloser){
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		if response.StatusCode == 404 {
			log.Fatal("There is no entry for this term!")
		} else {
			log.Fatalf("Status code error: %d %s", response.StatusCode, response.Status)
		}
	}

	return response.Body
}

func parseResponse(body io.ReadCloser) (entries []dict.Entry) {

	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal(err)
	}

	document.Find(".def-panel").Each(func(i int, s *goquery.Selection) {
		entries = append(entries, dict.Entry{
			Word:          s.Find(".word").Text(),
			Meaning:       s.Find(".meaning").Text(),
			Example:       s.Find(".example").Text(),
			Contributor:   s.Find(".contributor").Text(),
		})
	})

	return entries
}

func GetEntryForTerm(term string) {
	entries := parseResponse(getResponseForURL("https://www.urbandictionary.com/define.php?term=" + term))
	entry := entries[0]

	blue := color.New(color.FgBlue)
	wordBg := blue.Add(color.BgWhite)


	wordBg.Print(entry.Word)
	fmt.Println("\n")
	fmt.Println(entry.Meaning, "\n")
	fmt.Println(entry.Example, "\n")
	fmt.Println(entry.Contributor, "\n")
}
