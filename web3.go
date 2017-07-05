package web3

import (
	"github.com/regcostajr/go-web3/eth"
	"github.com/regcostajr/go-web3/personal"
	"github.com/regcostajr/go-web3/providers"
)

const Coin int64 = 1000000000000000000

type Web3 struct {
	provider providers.ProviderInterface
	Eth      *eth.Eth
	Personal *personal.Personal
}

func NewWeb3(provider providers.ProviderInterface) *Web3 {
	web3 := new(Web3)
	web3.Eth = eth.NewEth(provider)
	web3.Personal = personal.NewPersonal(provider)
	web3.provider = provider
	return web3
}
