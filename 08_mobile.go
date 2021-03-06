//program to parse html for mobile number

package main

import (
				"os"
				//"buffer"
				"fmt"
        //"strconv"
				//"io/ioutil"
        "strings"
				//"strconv"
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
	f, err := os.OpenFile("Mobile1.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// parse the contact info from html
	//doc.Find("p.contact-info span.mobilesv").Each(func(i int, s *goquery.Selection) {
    //fmt.Println(s.Text())
	//})
	//contact info ref
	m := map[string]string{
		"ba" : "-",
		"dc" : "+",
		"fe" : "(",
		"hg" : ")",
		"acb"  : "0",
		"yz" : "1",
		"wx" : "2",
		"vu" : "3",
		"ts" : "4",
		"rq" : "5",
		"po" : "6",
		"nm" : "7",
		"lk" : "8",
		"ji" : "9",
	}

	// parse contact info
	doc.Find("p.contact-info").Each(func(i int, s *goquery.Selection) {
		urlid, _ := s.Attr("onclick")

		//reset contactno
		contactno := ""

		s.Find("span.mobilesv").Each(func(i int, s *goquery.Selection) {

		//fmt.Println(s.Find("span.mobilesv").Attr("class"))
		//fmt.Println(s.Html())
			classname, _ := s.Attr("class")

			//build contactno
			contactno = contactno + m[classname[strings.Index(classname,"-") + 1 :len(classname)]]

			//fmt.Println(urlid + classname)

	  })
		//write urlid and contactno
		fmt.Fprintln(f, urlid[strings.Index(urlid,"044PXX44"):len(urlid) - 3] + `,` + contactno)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Fprintln(f, urlid[strings.Index(urlid,"044PXX44"):len(urlid) - 3] + `,` + classname[strings.Index(classname,"-") + 1 :len(classname)] + `,` + strconv.Itoa(i+1) )
		// if err != nil {
		// 	fmt.Println(err)
		// }
	})

	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
