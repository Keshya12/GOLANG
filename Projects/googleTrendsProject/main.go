package main

//Import Statements
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Defining Struct
// Variable of type struct
type Item struct {
	Title     string `xml:"title"`          // Exported field (starts with an uppercase letter)
	Traffic   string `xml:"approx_traffic"` // Exported field
	Link      string `xml:"link"`           // Exported field
	NewsItems []News `xml:"news_item"`      // Exported field
}

type News struct {
	HeadLine     string `xml:"news_item_title"` // Exported field
	HeadlineLink string `xml:"news_item_url"`   // Exported field
}

type rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"` // Exported field
	ItemList []Item `xml:"item"`  // Exported field
}

// Defining all Functions
func main() {
	var r rss
	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("**********************************************")
	fmt.Println("                GOOGLE TRENDS                 ")
	fmt.Println("**********************************************")
	for i := range r.Channel.ItemList {
		rank := (i + 1) //to start from 1 not 0
		fmt.Println("#           ", rank)
		fmt.Println("SEARCH TERM:", r.Channel.ItemList[i].Title)
		fmt.Println("LINK       :", r.Channel.ItemList[i].Link)

		for i = range r.Channel.ItemList[i].NewsItems {
			fmt.Println("HEADLINE   :", r.Channel.ItemList[i].NewsItems[i].HeadLine)
			fmt.Println("LINK TO ARTICLE:", r.Channel.ItemList[i].NewsItems[i].HeadLine)
			fmt.Println("****************************************************************")
		}
	}
}

func readGoogleTrends() []byte {
	response := getGoogleTrends()
	data, err := ioutil.ReadAll(response.Body) //to read the response to make the golang understand the response body
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return data
}

func getGoogleTrends() *http.Response {
	response, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=GB")
	//In golang if we use any http req or db the err should be handles in the next line
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return response
}
