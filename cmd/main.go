package main

import (
	"log"

	project "github.com/Chazarov/rest-app"
)

func main() {
	srv := new(project.Server)
	if err := srv.Run("8880"); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

}
