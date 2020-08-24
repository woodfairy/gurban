package core

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gurban/src/dict"
	"io"
	"log"
	"net/http"
)


func getResponseForURL(url string) (resp io.ReadCloser){
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", response.StatusCode, response.Status)
	}

	return response.Body
}

func parseResponse(body io.ReadCloser) (entries []dict.Entry) {
	/*
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("%s\n", data)*/
	document, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	document.Find(".def-panel").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		entries = append(entries, dict.Entry{
			Word:          s.Find(".word").Text(),
			Pronunciation: s.Find(".meaning").Text(),
			Meaning:       s.Find(".meaning").Text(),
		})


	})
	/*

	matcher := func(node *html.Node) (keep bool, exit bool) {

		for i, attr := range node.Attr {
			fmt.Println("[*] " + strconv.Itoa(i) + attr.Key + " - " + attr.Val)
			if attr.Key == "0class" {
				fmt.Println("[*] " + strconv.Itoa(i) + attr.Key + " - " + attr.Val)
				switch attr.Val {
				case "word":
					keep = true
				case "meaning":
					keep = true
					
				}
			}
		}
		
		return
	}

	nodes := TraverseNode(document, matcher)

	for i, node := range nodes {
		fmt.Println(i, renderNode(node))
	}

*/
	return entries
}

func GetEntryForTerm(term string) {
	entries := parseResponse(getResponseForURL("https://www.urbandictionary.com/define.php?term=" + term))
	entry := entries[0]
	fmt.Println(entry.Word, "\n")
	fmt.Println(entry.Meaning, "\n")
}
