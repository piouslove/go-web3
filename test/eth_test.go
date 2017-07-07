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
 * @file eth_test.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */
package test

import (
	"testing"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/eth/block"
	"github.com/regcostajr/go-web3/providers"
)

var ethClient = web3.NewWeb3(providers.NewHTTPProvider("http://127.0.0.1:8545", 100))

func TestGetBalance(t *testing.T) {

	balance, err := ethClient.Eth.GetBalance("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", block.LATEST)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(balance)

}

func TestSendTransaction(t *testing.T) {

	unlocked, err := ethClient.Personal.UnlockAccount("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", "tcpip01", 10)

	if !unlocked {
		t.Error(err)
		return
	}

	transaction, err := ethClient.Eth.SendTransaction("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", "0x007C0C7a7aB4aeDEE49bAb4c3a7dBfF9C675edCA", 1*web3.Coin, "0x7472616e73616374696f6e2073656e742062792074686520676f2d7765623320617069")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(transaction)

	search, err := ethClient.Eth.GetTransactionByHash("0xc5d2d3eb8c4463a6d42024bc975c2f70ae47e8d252897b237faaa0ecc29194ac")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(search.Gas)

}
