package BtcPrice

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

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
