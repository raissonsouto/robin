package main

import (
	"Robin/src"
	_ "Robin/src"
)

func main() {

	url, qtdRoutines, pathSize, _ := src.InitCli()
	 
	workBalancer := src.WorkBalancer(src.ValidChars, qtdRoutines)

	for i := 0; i < qtdRoutines; i++ {
		initialString := <-workBalancer

		src.Wg.Add(1)
		go src.Fuzz(url, initialString, pathSize, src.ValidChars)
	}
}
