package main

import (
	"eternal-infer-worker/chains/base"
	"eternal-infer-worker/config"
	"fmt"
	_ "net/http/pprof"
)

func main() {
	cnf, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	b, err := base.NewBaseChain(cnf)
	if err != nil {
		panic(err)
	}

	isStake, _ := b.IsStaked()
	fmt.Println(isStake)

	_ = b

}
