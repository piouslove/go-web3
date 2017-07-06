# Ethereum Go Client

You need to run a local Ethereum node to use this library.

This is a Ethereum compatible Go Client
which implements the 
[Eth JSON RPC Module](https://github.com/ethereum/wiki/wiki/JSON-RPC),
[Personal JSON RPC Module](https://github.com/paritytech/parity/wiki/JSONRPC-personal-module) and
[NET JSON RPC Module](https://github.com/paritytech/parity/wiki/JSONRPC-net-module#net_version).

## Status

This package is currently under active development. It is not already stable and the infrastructure is not complete and there are still several RPCs left to implement and the API is not stable yet.

## Installation

### go get

```bash
go get -u github.com/regcostajr/go-web3
```

### glide

```bash
glide get github.com/regcostajr/go-web3
```

## Usage

```go
import (
	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/eth/block"
	"github.com/regcostajr/go-web3/providers"
)

func test() {

	web3Client := web3.NewWeb3(providers.NewHTTPProvider("http://127.0.0.1:8545", 10))
	balance, err := web3Client.Eth.GetBalance("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", block.LATEST)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(balance)

}
```

More samples in the 'examples' directory

## Contribute!

### Requirements

* go 1.8.3

[Go installation instructions.](https://golang.org/doc/install)

## Testing

### Inside 'test' directory:
```bash
go test -v
```

### Or:
```bash
go test "file_name" -v
```

## License

Package go-web3 is licensed under the [GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html) License.
