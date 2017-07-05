package block

import "fmt"

func BLOCKNUMBER(blocknumber int) string {
	return fmt.Sprintf("%x", blocknumber)
}

const (
	// EARLIEST - Earliest block
	EARLIEST string = "earliest"
	// LATEST - latest block
	LATEST string = "latest"
	// PENDING - Pending block
	PENDING string = "pending"
)
