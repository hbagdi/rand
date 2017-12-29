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
	app.Commands = []cli.Command{
		{
			Name:  "name",
			Usage: "generate a random (first and last) name of person",
			Action: func(c *cli.Context) error {
				fmt.Println(gofakeit.Name())
				return nil
			},
		},
		{
			Name:  "email",
			Usage: "generate a random e-mail address",
			Action: func(c *cli.Context) error {
				fmt.Println(gofakeit.Email())
				return nil
			},
		},
		{
			Name:  "cell",
			Usage: "generate a random cellular number",
			Action: func(c *cli.Context) error {
				fmt.Println(gofakeit.Phone())
				return nil
			},
		},
		{
			Name:  "cc",
			Usage: "generate a random credit card number",
			Action: func(c *cli.Context) error {
				fmt.Println(gofakeit.CreditCardNumber())
				return nil
			},
		},
		{
			Name:  "ssn",
			Usage: "generate a random US Social Security Number (SSN)",
			Action: func(c *cli.Context) error {
				fmt.Println(gofakeit.SSN())
				return nil
			},
		},
	}
	app.Run(os.Args)
}
