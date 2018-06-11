package BtcPrice

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
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
