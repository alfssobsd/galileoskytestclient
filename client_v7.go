package main

import (
	"flag"
	"github.com/alfssobsd/galileoskytestclient/usecases"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("Test Client GalileoSky 7.0")
	selectedUseCase := flag.String("usecase", "onepack", "usecase for run 'movement' or 'onepack' (default:'onepack')")
	host := flag.String("host", "127.0.0.1", "host address (default: 127.0.0.1)")
	port := flag.Int("port", 9998, "host port (default:9998)")

	oneHexPack := flag.String("pack", "", "pack [usecase = onepack]")
	oneHexResponsePack := flag.String("response", "", "response pack [usecase = onepack]")

	movementConfigPath := flag.String("movement_config_path", "client_v7_test_movement.yml",
		"path to config [usecase = movement] (default: client_v7_test_movement.yml)")

	flag.Parse()

	if *selectedUseCase == "movement" {
		usecases.EmulateMovmentUseCase(*host, *port, *movementConfigPath)
	} else {
		usecases.SendOnePackUseCase(*host, *port,
			*oneHexPack, *oneHexResponsePack)
	}

}
