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
        "log"
        "net/http"
				"database/sql"
				"github.com/PuerkitoBio/goquery"
				"github.com/mattn/go-sqlite3"
        //"encoding/json"
        //"github.com/tidwall/gjson"
//				"github.com/gocolly/colly"
)

func main () {

	db, err := sql.Open("sqlite3", "C:\Users\264166\go\src\github.com\vivekanandaraman\broapp\listing.db")
	checkErr(err)

	// If the log file doesn't exist, create it or append to the file
	logfile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
			fmt.Println(err)
	}

	log.SetOutput(logfile)

	// error func defn
	func checkErr(err error) {
		if err != nil {
				log.Fatal(err)
		}

  client := &http.Client{}

  req, err := http.NewRequest("GET", "https://www.justdial.com/Chennai/Grocery-Stores-in-Selaiyur/nct-10237947/page-1", nil)
	checkErr(err)

  req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

  resp, err := client.Do(req)
	checkErr(err)

  defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	//create file if not already exists
	f, err := os.OpenFile("Mobile1.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

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
	  })
		//write urlid and contactno
		fmt.Fprintln(f, urlid[strings.Index(urlid,"044PXX44"):len(urlid) - 3] + `,` + contactno)

		// insert
		stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
		checkErr(err)

		res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
		checkErr(err)

	})
	err = f.Close()
	checkErr(err)
	err = logfile.Close()
	checkErr(err)
}
