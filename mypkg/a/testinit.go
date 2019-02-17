package a

import (
	"fmt"
	"iotestgo/mypkg/b"
)

func init() {
	fmt.Println("a package init")
}

func Testinita() {
	b.Testinitb()
}
