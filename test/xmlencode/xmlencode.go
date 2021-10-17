package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Results struct {
	Shops []Shop `xml:"shop"`
}

type Shop struct {
	Name string `xml:"name"`
}

func main() {
	shops := []Shop{
		{Name: "hoge"},
		{Name: "fuga"},
	}
	results := Results{Shops: shops}
	encoded, err := xml.Marshal(&results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encoded))
}
