package types

import (
	"fmt"
	"strconv"
	"strings"
)

type ComplexIntParameter int64

func (s ComplexIntParameter) ToHex() string {

	return fmt.Sprintf("0x%x", s)

}

type ComplexIntResponse string

func (s ComplexIntResponse) ToUInt64() uint64 {

	stringValue := string(s)

	cleaned := strings.Replace(stringValue, "0x", "", -1)
	sResult, _ := strconv.ParseUint(cleaned, 16, 64)

	return sResult

}

func (s ComplexIntResponse) ToInt64() int64 {

	stringValue := string(s)

	cleaned := strings.Replace(stringValue, "0x", "", -1)
	sResult, _ := strconv.ParseInt(cleaned, 16, 64)

	return sResult

}
