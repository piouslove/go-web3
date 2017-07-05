package test

import (
	"testing"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/eth/block"
	"github.com/regcostajr/go-web3/providers"
)

func Test(t *testing.T) {

	web3Client := web3.NewWeb3(providers.NewIPCProvider("/home/reginaldo/.local/share/io.parity.ethereum/jsonrpc.ipc"))
	balance, err := web3Client.Eth.GetBalance("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", block.LATEST)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(balance)

}
