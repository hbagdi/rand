package main

import (
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit"
	"github.com/urfave/cli"
)

func init() {
	gofakeit.Seed(0)
}

func main() {
	app := cli.NewApp()
	app.Name = "rand"
	app.Usage = "generate semantic psuedo-random data"
	app.Action = func(c *cli.Context) error {
		fmt.Println(gofakeit.Name())
		return nil
	}
	app.Run(os.Args)
}
