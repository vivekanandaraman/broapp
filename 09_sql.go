//program to parse html for mobile number

package main

import (
				//"os"
				//"buffer"
				//"fmt"
        //"strconv"
				//"io/ioutil"
        //"strings"
				//"strconv"
        //"log"
        //"net/http"
				"database/sql"
				//"github.com/PuerkitoBio/goquery"
				"github.com/mattn/go-sqlite3"
        //"encoding/json"
        //"github.com/tidwall/gjson"
//				"github.com/gocolly/colly"
)
func main () {
	db, err := sql.Open("sqlite3", "C:\Users\264166\go\src\github.com\vivekanandaraman\broapp\listing.db")
  Url := "044"
  Contact := "94454"
  db.exec("INSERT INTO Mobile1(UrlID, ContactNo) VALUES (?,?)",Url,Contact)
  db.close()
}
