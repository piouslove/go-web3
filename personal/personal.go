package personal

import (
	"fmt"

	"github.com/regcostajr/go-web3/dto"
	"github.com/regcostajr/go-web3/providers"
)

type Personal struct {
	provider providers.ProviderInterface
}

func NewPersonal(provider providers.ProviderInterface) *Personal {
	personal := new(Personal)
	personal.provider = provider
	return personal
}

func (personal *Personal) ListAccounts() ([]string, error) {

	pointer := &dto.RequestResult{}

	err := personal.provider.SendRequest(pointer, "personal_listAccounts", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToStringArray()

}

func (personal *Personal) NewAccount(password string) (string, error) {

	pointer := &dto.RequestResult{}

	err := personal.provider.SendRequest(&pointer, "personal_newAccount", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

func (personal *Personal) SendTransaction(from string, to string, value int64, hexData string) (string, error) {

	params := make([]dto.TransactionParameters, 1)

	params[0].From = from
	params[0].To = to
	params[0].Value = fmt.Sprintf("0x%x", value)
	params[0].Data = hexData

	pointer := &dto.RequestResult{}

	err := personal.provider.SendRequest(pointer, "personal_sendTransaction", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// UnlockAccount ...
func (personal *Personal) UnlockAccount(address string, password string, duration int) (bool, error) {

	params := make([]string, 3)
	params[0] = address
	params[1] = password
	params[2] = fmt.Sprintf("0x%x", duration)

	pointer := &dto.RequestResult{}

	err := personal.provider.SendRequest(pointer, "personal_unlockAccount", params)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}
