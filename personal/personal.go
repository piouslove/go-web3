package personal

import (
	"fmt"

	"github.com/regcostajr/go-web3/eth/dto"
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

// UnlockAccount ...
func (personal *Personal) UnlockAccount(address string, password string, duration int) (bool, error) {

	params := make([]string, 3)
	params[0] = address
	params[1] = password
	params[2] = fmt.Sprintf("0x%x", duration)

	unlock := dto.RequestResult{}
	pointer := &unlock

	err := personal.provider.SendRequest(pointer, "personal_unlockAccount", params)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}
