package main

import (
	"errors"
	"fmt"
	password "gebes.io/Password/src"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	password.Init()
	app := cli.NewApp()

	app.Name = "Password Generator"
	app.Usage = "Let's you quickly generate passwords without the usage of the web"

	flags := []cli.Flag{
		cli.IntFlag{
			Name:  "length",
			Value: 24,
			Usage: "Sets the length of the password",
		},
		cli.BoolFlag{
			Name:  "copy",
			Usage: "Defines if the password should be copied to the clipboard",
		},
		cli.BoolFlag{
			Name:  "no-print",
			Usage: "Prevents the password from being printed",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "Generate a new secure random password",
			Flags: flags,
			Action: func(c *cli.Context) error {

				length := c.Int("length")
				if length <= 0 {
					return errors.New("password length needs to be at least 1")
				}

				copy := c.Bool("copy")
				noPrint := c.Bool("no-print")

				password, err := password.GenerateRandomString(length)
				if err != nil {
					return errors.New("could not generate password")
				}

				if !noPrint {
					fmt.Println(password)
				}

				if copy {
					err := clipboard.WriteAll(password)
					if err != nil {
						return errors.New("could not copy to clipboard")
					}
					fmt.Println("Copied to clipboard!")
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
