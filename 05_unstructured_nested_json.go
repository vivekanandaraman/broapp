//program to parse nested json
//have to resolve errors
package main

import (
//        "io/ioutil"
				"os"
				"fmt"
        //"log"
        "net/http"
				"github.com/PuerkitoBio/goquery"
        "encoding/json"
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

        // parse jsonData
        var list1 []map[string]interface{}
        if err := json.Unmarshal([]byte(jsonData), &list1); err != nil {
            fmt.Println(err)
        }

				//create file
				f, err := os.Create("address.txt")
				if err != nil {
					fmt.Println(err)
				}

				line := ""
				//parse the array elements
        for i := range list1 {
					//parse individual map
          for k,v := range list1[i] {
            switch v := v.(type) {
              case string:
                fmt.Println(k, v, "(string)")
								switch k {
									case "name":
										line = v
									case "image":
										line =	line + v
								}
							case map[string]interface{}:
								for i, u := range v {
									switch uu := u.(type) {
										case "string":
											switch i {
												case "streetAddress":
													line = line + u
												case "addresslocality":
													line = line + u
												case "postalCode":
													line = line + u
												case "addressCountry":
													line = line + u
											}
									}
								}

              default:
                fmt.Println(k, v, "(unknown)")
            }
          }
					fmt.Fprintln(f, line)
					if err != nil {
						fmt.Println(err)
					}
        }
				err = f.Close()
				if err != nil {
					fmt.Println(err)
				}
}
