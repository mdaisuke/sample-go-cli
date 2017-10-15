package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("hello friend, %s!\n", c.Args().Get(0))
		return nil
	}

	app.Run(os.Args)
}
