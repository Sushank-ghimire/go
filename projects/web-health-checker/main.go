package main

import (
	"fmt"
	"log"
	"os"
	"web-health-checker/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("--------Website Health Checker------------")
	app := &cli.App{
		Name:  "web-health-checker",
		Usage: "Used to check the health of a site that is down or not",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "To check the provided domain is down or not",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Port number to check",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("p")
			if c.String("port") == "" {
				port = "80"
			}

			status := checker.Check(c.String("domain"), port)
			fmt.Println("Status : ", status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error occured : ", err)
		log.Fatal(err)
	}
}
