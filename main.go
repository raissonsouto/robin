package main

import (
	"Robin/src"
	_ "Robin/src"
)

func main() {

	src.InitCli()
	src.PrintInitialMessage()

	src.Wg.Add(1)
	go src.AlphabetGenerator()

	for i := 0; i < src.QtdRoutines; i++ {
		src.Wg.Add(1)
		go src.Fuzz()
	}

	src.Wg.Wait()
}
