//program to parse html for mobile number

package main

import (
				//"os"
				//"buffer"
				"fmt"
        //"strconv"
				//"io/ioutil"
        //"strings"
				//"strconv"
        //"log"
        //"net/http"
				"database/sql"
				//"github.com/PuerkitoBio/goquery"
			_"github.com/mattn/go-sqlite3"
        //"encoding/json"
        //"github.com/tidwall/gjson"
//				"github.com/gocolly/colly"
)
func main () {
	db, _ := sql.Open("sqlite3", "/Users/vivekanandaraman/go/src/github.com/vivekanandaraman/broapp/listing.db")
  Url := "044"
  Contact := "94454"
  stmt, _ := db.Prepare("INSERT INTO Mobile1(UrlID, ContactNo) VALUES (?,?)")
	stmt.Exec(Url,Contact)
	row, _ := db.Query("SELECT * FROM Mobile1")
	for row.Next() { // Iterate and fetch the records from result cursor
		var UrlId1 string
		var Contactno1 string
		row.Scan(&UrlId1, &Contactno1)
		fmt.Println(UrlId1, " " , Contactno1)
	}
	row.Close()
  db.Close()
}
