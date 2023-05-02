package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "got",
		Usage: "fight the loneliness!",
		Action: func(cCtx *cli.Context) error {
			cmd := exec.Command("git", "clone", "https://github.com/yuyinws/yuyinws")
			_, err := cmd.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println("clone end")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
