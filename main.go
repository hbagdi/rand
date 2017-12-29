// Copyright (c) 2017 Harry Bagdi
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

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
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Value: 1,
			Name:  "count, c",
			Usage: "repeat the command `n` times",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "name",
			Usage: "generate a random (first and last) name of person",
			Action: func(c *cli.Context) error {
				for i := 0; i < c.GlobalInt("count"); i++ {
					fmt.Println(gofakeit.Name())
				}
				return nil
			},
		},
		{
			Name:  "email",
			Usage: "generate a random e-mail address",
			Action: func(c *cli.Context) error {
				for i := 0; i < c.GlobalInt("count"); i++ {
					fmt.Println(gofakeit.Email())
				}
				return nil
			},
		},
		{
			Name:  "cell",
			Usage: "generate a random cellular number",
			Action: func(c *cli.Context) error {
				for i := 0; i < c.GlobalInt("count"); i++ {
					fmt.Println(gofakeit.Phone())
				}
				return nil
			},
		},
		{
			Name:  "cc",
			Usage: "generate a random credit card number",
			Action: func(c *cli.Context) error {
				for i := 0; i < c.GlobalInt("count"); i++ {
					fmt.Println(gofakeit.CreditCardNumber())
				}
				return nil
			},
		},
		{
			Name:  "ssn",
			Usage: "generate a random US Social Security Number (SSN)",
			Action: func(c *cli.Context) error {
				for i := 0; i < c.GlobalInt("count"); i++ {
					fmt.Println(gofakeit.SSN())
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
