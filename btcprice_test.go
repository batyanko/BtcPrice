package BtcPrice

import (
	"testing"
	"net/http"
	"fmt"
	"reflect"
)

func TestResponse(t *testing.T) {
	b := BtcPriceThing{}

	exampleResp, err := http.Get(priceRefUrl)
	b.GetResponse()
	if err != nil {
		fmt.Println("Could not get response from website.")
	}
	//Fails anyway
	if !reflect.DeepEqual(exampleResp, b.resp) {
		t.Errorf("did not get expected response")
	}

}