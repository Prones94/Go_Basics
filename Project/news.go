package main

import ("io/ioutil"
		"net/http"
		"fmt"
		"encoding/xml")

type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
		
func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations {
		fmt.Printf("\n%s",Location)
	}
}