package dto

import (
	"errors"
	"strconv"
	"strings"

	"github.com/regcostajr/go-web3/constants"

	"encoding/json"
)

type RequestResult struct {
	ID      int         `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *Error      `json:"error,omitempty"`
	Data    string      `json:"data,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (pointer *RequestResult) ToStringArray() ([]string, error) {

	result := (pointer).Result.([]interface{})

	if pointer.Error != nil {
		return nil, errors.New(pointer.Error.Message)
	}

	new := make([]string, len(result))
	for i, v := range result {
		new[i] = v.(string)
	}

	return new, nil

}

func (pointer *RequestResult) ToString() (string, error) {

	result := (pointer).Result.(interface{})

	if pointer.Error != nil {
		return "", errors.New(pointer.Error.Message)
	}

	return result.(string), nil

}

func (pointer *RequestResult) ToInt() (uint64, error) {

	result := (pointer).Result.(interface{})

	if pointer.Error != nil {
		return 0, errors.New(pointer.Error.Message)
	}

	hex := result.(string)

	cleaned := strings.Replace(hex, "0x", "", -1)

	numericResult, err := strconv.ParseUint(cleaned, 16, 64)

	return uint64(numericResult), err

}

func (pointer *RequestResult) ToBoolean() (bool, error) {

	result := (pointer).Result.(interface{})

	if pointer.Error != nil {
		return false, errors.New(pointer.Error.Message)
	}

	return result.(bool), nil

}

func (pointer *RequestResult) ToTransactionResponse() (*TransactionResponse, error) {

	result := (pointer).Result.(map[string]interface{})

	if pointer.Error != nil {
		return nil, errors.New(pointer.Error.Message)
	}

	if len(result) == 0 {
		return nil, errors.New(constants.NOTFOUND)
	}

	transactionResponse := &TransactionResponse{}

	marshal, err := json.Marshal(result)

	if err != nil {
		return nil, errors.New(constants.UNPARSEABLE)
	}

	json.Unmarshal([]byte(marshal), transactionResponse)

	return transactionResponse, nil

}
