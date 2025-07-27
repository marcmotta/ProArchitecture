// cmd/proarchitecture/main.go
package main

import (
	"flag"
	"log"
	"os"

	"proarchitecture/internal/proarchitecture"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := proarchitecture.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
