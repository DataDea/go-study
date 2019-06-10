package a

import (
	"fmt"
	"go-study/mypkg/b"
)

func init() {
	fmt.Println("a package init")
}

func Testinita() {
	b.Testinitb()
}
