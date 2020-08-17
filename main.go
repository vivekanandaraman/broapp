package main

import (
//        "io/ioutil"
				"fmt"
        "log"
        "net/http"
				"github.com/PuerkitoBio/goquery"
//				"github.com/gocolly/colly"
)

func main() {
        client := &http.Client{}

        req, err := http.NewRequest("GET", "https://www.justdial.com/Chennai/Grocery-Stores-in-Selaiyur/nct-10237947/page-1", nil)
        if err != nil {
                log.Fatalln(err)
        }

        req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

        resp, err := client.Do(req)
        if err != nil {
                log.Fatalln(err)
        }

        defer resp.Body.Close()
        // body, err := ioutil.ReadAll(resp.Body)
        // if err != nil {
        //         log.Fatalln(err)
        // }
				//
        // log.Println(string(body))

				// Load the HTML document
				doc, err := goquery.NewDocumentFromReader(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}

				// Find the address items
				txt1 := doc.Find("script[type='application/ld+json']").Eq(3).Text()
        //doc.Find("script[type=application/ld+json']").Each(func(i int, s *goquery.Selection) {
					// For each item found, get the band and title
         	//txt1 := s.Find().Eq(i).Text()
				//	title := s.Find("i").Text()
				//fmt.Printf("Type of selector = %T",txt1)
				fmt.Println(txt1)
				//})
}
