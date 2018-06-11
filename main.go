package BtcPrice

import (
	"fmt"
	"net/http"
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




