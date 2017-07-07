/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file send-transaction.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package examples

import (
	"fmt"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
)

func SendTransactionSample() {

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

	transaction, err := web3Client.Eth.GetTransactionByHash(hash)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(transaction.Gas)

}
