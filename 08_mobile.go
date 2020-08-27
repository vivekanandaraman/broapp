//program to parse html for mobile number

package main

import (
				//"os"
				//"buffer"
				"fmt"
        //"strconv"
				//"io/ioutil"
        //"strings"
        //"log"
        "net/http"
				"github.com/PuerkitoBio/goquery"
        //"encoding/json"
        //"github.com/tidwall/gjson"
//				"github.com/gocolly/colly"
)

func main() {
  client := &http.Client{}

  req, err := http.NewRequest("GET", "https://www.justdial.com/Chennai/Grocery-Stores-in-Selaiyur/nct-10237947/page-1", nil)
  if err != nil {
    fmt.Println(err)
  }
	fmt.Println("hi")
  req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }

  defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// parse the contact info from html
	//doc.Find("p.contact-info span.mobilesv").Each(func(i int, s *goquery.Selection) {
    //fmt.Println(s.Text())
	//})
	doc.Find("p.contact-info").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Attr("onclick"))
		fmt.Println(s.Find("span").Attr("class"))
		//fmt.Println(s.Html())
	})

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(jsonData)
}
