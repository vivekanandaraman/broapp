package main

import (
//        "io/ioutil"
				"fmt"
        "log"
        "net/http"
				"github.com/PuerkitoBio/goquery"
        "encoding/json"
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

				// Load the HTML document
				doc, err := goquery.NewDocumentFromReader(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}

				// Find the address items
				jsonData := doc.Find("script[type='application/ld+json']").Eq(3).Text()
				fmt.Println(jsonData)

        // parse jsonData
        var list1 []map[string]interface{}
        if err := json.Unmarshal([]byte(jsonData), &list1); err != nil {
            panic(err)
        }
        for i := range list1 {
          fmt.Println(list1[i])
          for k,v := range list1[i] {
            switch v := v.(type) {
              case string:
                fmt.Println(k, v, "(string)")
              case float64:
                fmt.Println(k, v, "(float64)")
            }
          }
        }

}
