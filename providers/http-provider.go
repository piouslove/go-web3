package providers

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"encoding/json"

	"github.com/regcostajr/go-web3/providers/util"
)

type HTTPProvider struct {
	address string
	timeout int32
}

func NewHTTPProvider(address string, timeout int32) *HTTPProvider {
	provider := new(HTTPProvider)
	provider.address = address
	provider.timeout = timeout
	return provider
}

func (provider HTTPProvider) SendRequest(v interface{}, method string, params interface{}) error {

	bodyString := util.JSONRPCObject{Version: "2.0", Method: method, Params: params, ID: rand.Intn(100)}

	body := strings.NewReader(bodyString.AsJsonString())
	req, err := http.NewRequest("POST", provider.address, body)
	if err != nil {
		return err
	}
	/*req.Header.Set("Content-Type", "application/x-www-form-urlencoded")*/
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	var netClient = &http.Client{
		Timeout: time.Second * time.Duration(provider.timeout),
	}

	resp, err := netClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var bodyBytes []byte

	if resp.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(bodyBytes, v)

}
