//program to parse html for mobile number

package main

import (
				"os"
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

	//create file if not already exists
	f, err := os.OpenFile("Mobile.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// parse the contact info from html
	//doc.Find("p.contact-info span.mobilesv").Each(func(i int, s *goquery.Selection) {
    //fmt.Println(s.Text())
	//})
	doc.Find("p.contact-info").Each(func(i int, s *goquery.Selection) {
		urlid, _ := s.Attr("onclick")
		s.Find("span.mobilesv").Each(func(i int, s *goquery.Selection) {

		//fmt.Println(s.Find("span.mobilesv").Attr("class"))
		//fmt.Println(s.Html())
			classname, _ := s.Attr("class")
			//fmt.Println(urlid + classname)
			fmt.Fprintln(f, urlid + classname)
			if err != nil {
				fmt.Println(err)
			}
	  })
	})

	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
