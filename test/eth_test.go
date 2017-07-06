package test

import (
	"testing"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/eth/block"
	"github.com/regcostajr/go-web3/providers"
)

var web3Client = web3.NewWeb3(providers.NewHTTPProvider("http://127.0.0.1:8545", 100))

func TestGetBalance(t *testing.T) {

	balance, err := web3Client.Eth.GetBalance("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", block.LATEST)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(balance)

}

func TestSendTransaction(t *testing.T) {

	unlocked, err := web3Client.Personal.UnlockAccount("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", "password", 10)

	if !unlocked {
		t.Error(err)
		return
	}

	transaction, err := web3Client.Eth.SendTransaction("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", "0x007C0C7a7aB4aeDEE49bAb4c3a7dBfF9C675edCA", 1*web3.Coin, "0x")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(transaction)

	search, err := web3Client.Eth.GetTransactionByHash("0xc5d2d3eb8c4463a6d42024bc975c2f70ae47e8d252897b237faaa0ecc29194ac")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(search.Gas)

}
