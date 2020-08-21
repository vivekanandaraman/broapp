//program to parse nested json
//gjson

//curl sample to get GeoLocation
// {"lil":"12.919598","lon":"80.1447211"}
// curl 'https://www.justdial.com/functions/maps.php' \
//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36' \
//   -H 'referer: https://www.justdial.com/Chennai/Rajendran-Supermarket-Near-Parasakthi-Nagar-Selaiyur/044PXX44-XX44-100814143544-H1X2_BZDET' \
//   --data-raw 'doc=044PXX44.XX44.100814143544.H1X2'

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

				// Find the url items
				jsonData := doc.Find("script[type='application/ld+json']").Eq(2).Text()
				fmt.Println(jsonData)

        //var str1 string
        //url
        var url string

				//create file if not already exists
				f, err := os.OpenFile("Geo1.txt",
					os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println(err)
				}
        //parse the json
        for i := 0; i < 10; i++ {
          url = "itemListElement." + strconv.Itoa(i) + ".url"
					//fmt.Println(url)
					value := gjson.Get(jsonData, url)
          //fmt.Println(value.Str)

					//Request to get GeoLocation
					client := &http.Client{}
	        req, err := http.NewRequest("GET", value.Str, nil)
	        if err != nil {
	                fmt.Println(err)
	        }

	        req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

					req.Header.Add("referer", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

	        resp, err := client.Do(req)
	        if err != nil {
	                fmt.Println(err)
	        }

	        defer resp.Body.Close()


					//write to file
					fmt.Fprintln(f, value.Str)
					if err != nil {
						fmt.Println(err)
					}
				}
				err = f.Close()
				if err != nil {
					fmt.Println(err)
				}
}
