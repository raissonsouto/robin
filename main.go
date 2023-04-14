package main

import (
	"Robin/src"
	_ "Robin/src"
)

func main() {

	url, qtdRoutines, pathSize, _, _, _ := src.InitCli()
	var initialString string

	prefixChan := src.GetPrefixes(qtdRoutines)

	for i := 0; i < qtdRoutines; i++ {
		initialString = <-prefixChan

		src.Wg.Add(1)
		exponent := pathSize - len(initialString)
		go src.Fuzz(url, exponent)
	}
}
