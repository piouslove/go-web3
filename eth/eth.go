package eth

import (
	"fmt"

	"github.com/regcostajr/go-web3/dto"
	"github.com/regcostajr/go-web3/providers"
)

type Eth struct {
	provider providers.ProviderInterface
}

func NewEth(provider providers.ProviderInterface) *Eth {
	eth := new(Eth)
	eth.provider = provider
	return eth
}

func (eth *Eth) EstimateGas(from string, to string, value int64, hexData string) (uint64, error) {

	params := make([]dto.TransactionParameters, 1)

	params[0].From = from
	params[0].To = to
	params[0].Value = fmt.Sprintf("0x%x", value)
	params[0].Data = hexData

	transaction := dto.RequestResult{}
	pointer := &transaction

	err := eth.provider.SendRequest(&transaction, "eth_estimateGas", params)

	if err != nil {
		return 0, err
	}

	return pointer.ToInt()

}

// GetBalance ...
func (eth *Eth) GetBalance(address string, blockNumber string) (uint64, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = blockNumber

	balance := dto.RequestResult{}
	pointer := &balance

	err := eth.provider.SendRequest(pointer, "eth_getBalance", params)

	if err != nil {
		return 0, err
	}

	return pointer.ToInt()

}

func (eth *Eth) GetTransactionByHash(hash string) (*dto.TransactionResponse, error) {

	params := make([]string, 1)
	params[0] = hash

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_getTransactionByHash", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToTransactionResponse()

}

func (eth *Eth) ListAccounts() ([]string, error) {

	accounts := dto.RequestResult{}
	pointer := &accounts

	err := eth.provider.SendRequest(pointer, "eth_accounts", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToStringArray()

}

func (eth *Eth) SendTransaction(from string, to string, value int64, hexData string) (string, error) {

	gasResult, err := eth.EstimateGas(from, to, value, hexData)

	if err != nil {
		return "", err
	}

	params := make([]dto.TransactionParameters, 1)

	params[0].From = from
	params[0].To = to
	params[0].Gas = fmt.Sprintf("0x%x", gasResult)
	params[0].GasPrice = "0x9184e72a000"
	params[0].Value = fmt.Sprintf("0x%x", value)
	params[0].Data = hexData

	transaction := dto.RequestResult{}
	pointer := &transaction

	err = eth.provider.SendRequest(&transaction, "eth_sendTransaction", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
