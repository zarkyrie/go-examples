package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	str := "ljh"

	decodeString := hex.EncodeToString([]byte(str))
	fmt.Println(decodeString)
}
