// chain project main.go
package main

import (
	"chain/model"
)

func main() {
	bc := model.NewBlockchain()
	bc.SendData("send 1 BTC to joe")
	bc.SendData("send 2 EOS to jovan")
	bc.Print()
}
