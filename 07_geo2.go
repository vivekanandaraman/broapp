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
				//"os"
				//"buffer"
				"fmt"
        "strconv"
				"io/ioutil"
        "strings"
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

        //parse the json
        for i := 0; i < 10; i++ {
          url = "itemListElement." + strconv.Itoa(i) + ".url"
					//fmt.Println(url)

					//parse json data to be passed to referrer
					value := gjson.Get(jsonData, url)
          fmt.Println(value.Str)

					//substr to get the post data
					postdata := "doc=" + strings.ReplaceAll(value.Str[strings.Index(value.Str,"044PXX44"):len(value.Str) - 6],"-",".")

					fmt.Println(postdata)

					//Request to get GeoLocation
					client := &http.Client{}
	        req, err := http.NewRequest("POST", "https://www.justdial.com/functions/maps.php", strings.NewReader(postdata))
	        if err != nil {
	                fmt.Println(err)
	        }

	        req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

					req.Header.Add("Referer", value.Str)

	        resp, err := client.Do(req)
	        if err != nil {
	                fmt.Println(err)
	        }

					Geojsoni, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Println(err)
					}

	        defer resp.Body.Close()

					Geojsons := string(Geojsoni)
					fmt.Printf(Geojsons)
				}
}
