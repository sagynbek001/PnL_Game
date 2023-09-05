package main

import (
	"log"

	"gitlab.com/laboct2021/pnl-game/external/rest"
)

const PathToConfig = "./config/config.yaml"

func main() {

	log.Fatal(rest.Run(PathToConfig))
}
