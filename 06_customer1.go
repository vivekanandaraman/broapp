//program to parse nested json
//gjson
package main

import (
//        "io/ioutil"
				"os"
				"fmt"
        "strconv"
        //"strings"
        //"log"
        "net/http"
				"github.com/PuerkitoBio/goquery"
        //"encoding/json"
        "github.com/tidwall/gjson"
//				"github.com/gocolly/colly"
)

func main() {
        client := &http.Client{}

        req, err := http.NewRequest("GET", "https://www.justdial.com/Chennai/Grocery-Stores-in-Selaiyur/nct-10237947/page-1", nil)
        if err != nil {
                fmt.Println(err)
        }

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

				// Find the address items
				jsonData := doc.Find("script[type='application/ld+json']").Eq(3).Text()
				//fmt.Println(jsonData)
        var str1 string
        //name,streetAddress,addresslocality,postalCode,addressCountry,image
        var name, streetAddress, addresslocality, postalCode, addressCountry, image string
				//create file if not already exists
				f, err := os.OpenFile("Customer1.txt",
					os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println(err)
				}
        //skip 0th element
        for i := 1; i < 11; i++ {
          name = strconv.Itoa(i)+ ".name"
          streetAddress = strconv.Itoa(i)+ ".address.streetAddress"
          addresslocality = strconv.Itoa(i)+ ".address.addresslocality"
          postalCode = strconv.Itoa(i)+ ".address.postalCode"
          addressCountry = strconv.Itoa(i)+ ".address.addressCountry"
          image = strconv.Itoa(i)+ ".image"

					value := gjson.GetMany(jsonData, name, streetAddress, addresslocality, postalCode, addressCountry, image)
					//fmt.Println(value)
          //last element shud not have comma, build csv to write to file
          str1=""
          for i,v := range value {
            if i == 5 { str1 = str1 + `"` + v.Str + `"`
            } else    { str1 = str1 + `"` + v.Str + `",` }
          }
          //fmt.Println(str1)
          fmt.Fprintln(f, str1)
          if err != nil {
            fmt.Println(err)
          }
        }
        err = f.Close()
        if err != nil {
          fmt.Println(err)
        }
}
