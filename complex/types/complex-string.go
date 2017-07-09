package types

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type ComplexString string

func (s ComplexString) ToHex() string {

	return fmt.Sprintf("0x%x", s)

}

func (s ComplexString) ToString() string {

	stringValue := string(s)

	cleaned := strings.Replace(stringValue, "0x", "", -1)
	sResult, _ := hex.DecodeString(cleaned)

	return string(sResult)

}
