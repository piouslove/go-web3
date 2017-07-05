package providers

import (
	"encoding/json"
	"math/rand"
	"net"
	"path/filepath"

	"github.com/regcostajr/go-web3/providers/util"
)

type IPCProvider struct {
	endpoint string
}

func NewIPCProvider(endpoint string) *IPCProvider {
	provider := new(IPCProvider)
	provider.endpoint, _ = filepath.Abs(endpoint)
	return provider
}

func (provider IPCProvider) SendRequest(v interface{}, method string, params interface{}) error {

	bodyString := util.JSONRPCObject{Version: "2.0", Method: method, Params: params, ID: rand.Intn(100)}

	client, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: provider.endpoint, Net: "unix"})
	defer client.Close()

	encoder := json.NewEncoder(client)
	err = encoder.Encode(bodyString)

	if err != nil {
		return err
	}

	err = json.NewDecoder(client).Decode(v)

	if err != nil {
		return err
	}

	return err

}
