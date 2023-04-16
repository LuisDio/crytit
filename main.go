package main

import (
	"fmt"

	"github.com/LuisDio/cryptit/decrypt"
	"github.com/LuisDio/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
