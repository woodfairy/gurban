package core

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"gurban/src/dict"
	"io"
	"log"
	"net/http"
	"strconv"
)


func getResponseForURL(url string) (resp *http.Response){
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func parseResponse(response *http.Response) (entry dict.Entry) {
	/*
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("%s\n", data)*/
	document, err := html.Parse(response.Body)

	if err != nil {
		log.Fatal(err)
	}

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


	return dict.Entry{
		Word:          "test",
		Pronunciation: "test",
		Meaning:       "test",
	}
}

func TraverseNode(
	doc *html.Node,
	matcher func(node *html.Node) (bool, bool)) (nodes []*html.Node) {

	var keep, exit bool
	var f func(*html.Node)
	f = func(n *html.Node) {
		keep, exit = matcher(n)
		if keep {
			nodes = append(nodes, n)
		}
		if exit {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return nodes
}

func GetEntryForTerm(term string) {
	entry := parseResponse(getResponseForURL("https://www.urbandictionary.com/define.php?term=" + term))
	fmt.Println(entry)
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
