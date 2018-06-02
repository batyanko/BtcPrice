package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

func main() {
	fmt.Println("halloo")

	resp, err := http.Get("https://blockchain.info/tobtc?currency=USD&value=1")
	//resp, err := http.Get("http://www.batyanko.com")
	body, err2 := ioutil.ReadAll(resp.Body)

	priceUsd, err3:= strconv.ParseFloat(string(body), 32)

	priceBtc := 1/priceUsd

	fmt.Println(priceBtc, err, err2, err3)

}

