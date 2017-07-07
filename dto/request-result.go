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
 * @file request-result.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

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

	if pointer.Error != nil {
		return nil, errors.New(pointer.Error.Message)
	}

	result := (pointer).Result.([]interface{})

	new := make([]string, len(result))
	for i, v := range result {
		new[i] = v.(string)
	}

	return new, nil

}

func (pointer *RequestResult) ToString() (string, error) {

	if pointer.Error != nil {
		return "", errors.New(pointer.Error.Message)
	}

	result := (pointer).Result.(interface{})

	return result.(string), nil

}

func (pointer *RequestResult) ToInt() (uint64, error) {

	if pointer.Error != nil {
		return 0, errors.New(pointer.Error.Message)
	}

	result := (pointer).Result.(interface{})

	hex := result.(string)

	cleaned := strings.Replace(hex, "0x", "", -1)

	numericResult, err := strconv.ParseUint(cleaned, 16, 64)

	return uint64(numericResult), err

}

func (pointer *RequestResult) ToBoolean() (bool, error) {

	if pointer.Error != nil {
		return false, errors.New(pointer.Error.Message)
	}

	result := (pointer).Result.(interface{})

	return result.(bool), nil

}

func (pointer *RequestResult) ToTransactionResponse() (*TransactionResponse, error) {

	if pointer.Error != nil {
		return nil, errors.New(pointer.Error.Message)
	}

	result := (pointer).Result.(map[string]interface{})

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
