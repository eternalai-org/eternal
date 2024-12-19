package main

import (
	"eternal-infer-worker/chains/base"
	_ "net/http/pprof"
)

func main() {
	b, err := base.NewBaseChain("https://mainnet.base.org")
	if err != nil {
		panic(err)
	}

	_ = b

}
