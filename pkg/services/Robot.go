package services

import (
	"fmt"
	"github.com/vazha/wsDataProvider/pkg/services/kraken"
)

type Exchange interface {
	Start()
	Stop()
	Subscribe([]string)
}

func Sync(e Exchange){
	fmt.Println(e.Stop)
}

func Engine(){
	k := kraken.New("wss://kraken.com")
	Sync(k)
}