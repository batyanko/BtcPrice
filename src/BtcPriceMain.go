package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"time"
)

type BtcPriceThing struct {
	priceBtc float64
}

func main() {
	fmt.Println("halloo")
	b := BtcPriceThing{0}
	ch := make(chan float64)

	go func() {
		for {
			b.getResponse(ch)
			time.Sleep(10 *time.Second)
		}
	}()
	for i := range ch {
		fmt.Println(i)
	}

}

func (b *BtcPriceThing) getResponse(ch chan float64) {
	resp, err := http.Get("https://blockchain.info/tobtc?currency=USD&value=1")
	if err != nil {
		fmt.Println("Could not get response from website.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response from website.")
	}

	priceUsd, err:= strconv.ParseFloat(string(body), 32)
	if err != nil {
		fmt.Println("Could not parse Btc price date.")
	}

	b.priceBtc = 1/priceUsd
	ch <- b.priceBtc
}