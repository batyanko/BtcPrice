package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"time"
)

type BtcPriceThing struct {
	//Price of BTC in USD
	priceBtc float64
	//Price of USD in BTC
	priceUsd float64
	resp *http.Response
	err error
	body []byte
}

const waitTime = 10
const priceRefUrl = "https://blockchain.info/tobtc?currency=USD&value=1"

func main() {
	fmt.Println("halloo")
	b := BtcPriceThing{}
	ch := make(chan float64)

	go func() {
		for {
			b.GetResponse()
			b.ReadResponse(b.resp)
			b.parseResponse(b.body)
			b.priceBtc = 1/b.priceUsd
			ch <- b.priceBtc
			time.Sleep(waitTime *time.Second)
		}
	}()

	for i := range ch {
		fmt.Println(i)
	}

}

func (b *BtcPriceThing) GetResponse() {
	b.resp, b.err = http.Get(priceRefUrl)
	if b.err != nil {
		fmt.Println("Could not get response from website.")
	}
}

func (b *BtcPriceThing) ReadResponse(resp *http.Response)  {

	b.body, b.err = ioutil.ReadAll(resp.Body)
	if b.err != nil {
		fmt.Println("Could not read response from website.")
	}
}

func (b *BtcPriceThing) parseResponse(body []byte) {
	b.priceUsd, b.err = strconv.ParseFloat(string(body), 32)
	if b.err != nil {
		fmt.Println("Could not parse Btc price date.")
	}
}



