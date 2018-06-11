package BtcPrice

import (
	"fmt"
	"time"
)

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




