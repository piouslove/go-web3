package examples

import (
	"fmt"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
)

func SendTransactionIPCProviderSample() {

	from := "0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42"
	to := "0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42"
	value := 1 * web3.Coin
	data := "transaction sent by the go-web3 api"

	web3Client := web3.NewWeb3(providers.NewIPCProvider("~/.local/share/io.parity.ethereum/jsonrpc.ipc"))
	hash, err := web3Client.Eth.SendTransaction(from, to, value, data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(hash)

}

func SendTransactionHTTPProviderSample() {

	from := "0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42"
	password := "from account password here"
	to := "0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42"
	value := 1 * web3.Coin
	data := "transaction sent by the go-web3 api"

	web3Client := web3.NewWeb3(providers.NewHTTPProvider("http://127.0.0.1:8545", 10))

	freedom, err := web3Client.Personal.UnlockAccount(from, password, 10)

	if !freedom {
		fmt.Println(err)
		return
	}

	hash, err := web3Client.Eth.SendTransaction(from, to, value, data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(hash)

}
