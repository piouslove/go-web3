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
 * @file eth.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package eth

import (
	"github.com/regcostajr/go-web3/complex/types"
	"github.com/regcostajr/go-web3/dto"
	"github.com/regcostajr/go-web3/providers"
)

// Eth - The Eth Module
type Eth struct {
	provider providers.ProviderInterface
}

// NewEth - Eth Module constructor to set the default provider
func NewEth(provider providers.ProviderInterface) *Eth {
	eth := new(Eth)
	eth.provider = provider
	return eth
}

// GetProtocolVersion - Returns the current ethereum protocol version.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_protocolversion
// Parameters:
//    - none
// Returns:
// 	  - String - The current ethereum protocol version
func (eth *Eth) GetProtocolVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_protocolVersion", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// IsSyncing - Returns an object with data about the sync status or false.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_syncing
// Parameters:
//    - none
// Returns:
// 	  - Object|Boolean, An object with sync status data or FALSE, when not syncing:
//    	- startingBlock: 	QUANTITY - The block at which the import started (will only be reset, after the sync reached his head)
//    	- currentBlock: 	QUANTITY - The current block, same as eth_blockNumber
//    	- highestBlock: 	QUANTITY - The estimated highest block
func (eth *Eth) IsSyncing() (*dto.SyncingResponse, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_syncing", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToSyncingResponse()

}

// EstimateGas - Makes a call or transaction, which won't be added to the blockchain and returns the used gas, which can be used for estimating the used gas.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_estimategas
// Parameters:
//    - See eth_call parameters, expect that all properties are optional. If no gas limit is specified geth uses the block gas limit from the pending block as an
// 		upper bound. As a result the returned estimate might not be enough to executed the call/transaction when the amount of gas is higher than the pending block gas limit.
// Returns:
//    - QUANTITY - the amount of gas used.
func (eth *Eth) EstimateGas(from string, to string, value types.ComplexIntParameter, hexData types.ComplexString) (int64, error) {

	params := make([]dto.TransactionParameters, 1)

	params[0].From = from
	params[0].To = to
	params[0].Value = value.ToHex()
	params[0].Data = hexData.ToHex()

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(&pointer, "eth_estimateGas", params)

	if err != nil {
		return 0, err
	}

	return pointer.ToInt()

}

// GetBalance - Returns the balance of the account of given address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getbalance
// Parameters:
//    - DATA, 20 Bytes - address to check for balance.
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
// Returns:
// 	  - QUANTITY - integer of the current balance in wei.
func (eth *Eth) GetBalance(address string, blockNumber string) (int64, error) {

	params := make([]string, 2)
	params[0] = string(address)
	params[1] = blockNumber

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_getBalance", params)

	if err != nil {
		return 0, err
	}

	return pointer.ToInt()

}

// GetBlockNumber - Returns the number of most recent block.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_blocknumber
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the current block number the client is on.
func (eth *Eth) GetBlockNumber() (int64, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_blockNumber", nil)

	if err != nil {
		return 0, err
	}

	return pointer.ToInt()

}

// GetTransactionByHash - Returns the information about a transaction requested by transaction hash.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash
// Parameters:
//    - DATA, 32 Bytes - hash of a transaction
// Returns:
//    1. Object - A transaction object, or null when no transaction was found
//    - hash: DATA, 32 Bytes - hash of the transaction.
//    - nonce: QUANTITY - the number of transactions made by the sender prior to this one.
//    - blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
//    - blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
//    - transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
//    - from: DATA, 20 Bytes - address of the sender.
//    - to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
//    - value: QUANTITY - value transferred in Wei.
//    - gasPrice: QUANTITY - gas price provided by the sender in Wei.
//    - gas: QUANTITY - gas provided by the sender.
//    - input: DATA - the data send along with the transaction.
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

// ListAccounts - Returns a list of addresses owned by client.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_accounts
// Parameters:
//    - none
// Returns:
//    - Array of DATA, 20 Bytes - addresses owned by the client.
func (eth *Eth) ListAccounts() ([]string, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_accounts", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToStringArray()

}

// SendTransaction - Creates new message call transaction or a contract creation, if the data field contains code.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
// Parameters:
//    1. Object - The transaction object
//    - from: 		DATA, 20 Bytes - The address the transaction is send from.
//    - to: 		DATA, 20 Bytes - (optional when creating new contract) The address the transaction is directed to.
//    - gas: 		QUANTITY - (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
//    - gasPrice: 	QUANTITY - (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas
//    - value: 		QUANTITY - (optional) Integer of the value send with this transaction
//    - data: 		DATA - The compiled code of a contract OR the hash of the invoked method signature and encoded parameters. For details see Ethereum Contract ABI
//    - nonce: 		QUANTITY - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
// Returns:
//	  - DATA, 32 Bytes - the transaction hash, or the zero hash if the transaction is not yet available.
// Use eth_getTransactionReceipt to get the contract address, after the transaction was mined, when you created a contract.
func (eth *Eth) SendTransaction(from string, to string, value types.ComplexIntParameter, hexData types.ComplexString) (string, error) {

	params := make([]dto.TransactionParameters, 1)

	params[0].From = from
	params[0].To = to
	params[0].Value = value.ToHex()
	params[0].Data = hexData.ToHex()

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(&pointer, "eth_sendTransaction", params)

	if err != nil {
		return "", err
	}

	response, err := pointer.ToString()

	return response, err

}
