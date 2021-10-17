package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

// type Plant struct {
// 	// Id      int      `xml:"id,attr"`
// 	shopName    string   `xml:"name"`
// 	Origin  []string `xml:"origin"`
// }
type Schema struct {
	Results Results `xml:"results"`
}

type Results struct {
	Shops []Shop `xml:"shop"`
}

type Shop struct {
	Name string `xml:"name"`
}

func main() {
	content, err := ioutil.ReadFile("sample.xml")
	if err != nil {
		log.Fatal(err)
	}
	var p Schema
	if err := xml.Unmarshal(content, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
}
