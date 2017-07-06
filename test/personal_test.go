package test

import (
	"testing"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
)

var personalClient = web3.NewWeb3(providers.NewHTTPProvider("http://127.0.0.1:8545", 100))

func TestListAccounts(t *testing.T) {

	list, err := personalClient.Personal.ListAccounts()

	if err != nil {
		t.Error(err)
		return
	}

	for index := 0; index < len(list); index++ {
		t.Log(list[index])
	}
}
