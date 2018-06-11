package BtcPrice

import (
	"testing"
	"net/url"
)

func TestUrl(t *testing.T) {
	_, err := url.ParseRequestURI(priceRefUrl)
	if err != nil {
		t.Errorf("url \"%s\" is incorrect, it could not be parsed. Error: %s", priceRefUrl, err)
	}
}
