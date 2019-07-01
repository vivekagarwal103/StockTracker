package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	//	"io/ioutil"
	"log"
	"net/http"
	//	"strings"
)

func main() {
	// Make request
	response, err := http.Get("https://nseindia.com/live_market/dynaContent/live_watch/get_quote/GetQuote.jsp?symbol=INFY")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create output file
	//outFile, err := os.Create("output.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//	defer outFile.Close()
	//_, err = io.Copy(outFile, response.Body)
	//if err != nil {
	//		log.Fatal(err)
	//	}
	/*	z := html.NewTokenizer(bufio.NewReader(response.Body))
		for {
			tt := z.Next()
			switch tt {
			case html.ErrorToken:
				return
			case html.StartTagToken:
				t := z.Token()
				switch t.Data {
				case "div":
					//				fmt.Println(t.Data)
					if strings.Contains(t.String(), "responseDiv") {
						fmt.Println(t.String())
						fmt.Println("t.Attr[0].Key : " + t.Attr[0].Key + "t.Attr[0].Val :" + t.Attr[0].Val)
						fmt.Println("t.Attr[1].Key : " + t.Attr[1].Key + "t.Attr[1].Val :" + t.Attr[1].Val)
						}
				}
			}
		}*/
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	//	fmt.Println("\n\n", bodyString, "\n\n")
	// Find a substr
	titleStartIndex := strings.Index(bodyString, "<div id=\"responseDiv\" style=\"display:none\">")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	titleStartIndex += len("<div id=\"responseDiv\" style=\"display:none\">") + 22
	// Find the index of the closing tag
	titleEndIndex := strings.Index(bodyString, "<!-- Content Big Start -->")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}
	titleEndIndex = titleEndIndex - len("</div>") - 9

	// (Optional)
	// Copy the substring in to a separate variable so the
	// variables with the full document data can be garbage collected
	pageTitle := []byte(bodyString[titleStartIndex:titleEndIndex])

	// Print out the result
	fmt.Printf("Page title: %s\n", pageTitle)
	// Copy data from HTTP response to file

}
