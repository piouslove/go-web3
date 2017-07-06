package examples

import (
	"fmt"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
)

func SendTransactionHTTPProviderSample() {

	from := "0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42"
	password := "from account password here"
	to := "0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42"
	value := 1 * web3.Coin
	data := "0x7472616e73616374696f6e2073656e742062792074686520676f2d7765623320617069"

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
