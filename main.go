package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("error")
	}
	APIkey := os.Args[1]
	small_area := os.Args[2]

	err := GourmetSerach(APIkey, small_area)
	if err != nil {
		log.Fatal(err)
	}

}

func GourmetSerach(APIkey, small_area string) error {
	url := fmt.Sprintf("https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?key=%s&small_area=%s", APIkey, small_area)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
