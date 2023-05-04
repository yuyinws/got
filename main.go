package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "got",
		Usage: "Fast project scaffolding",
		Action: func(ctx *cli.Context) error {
			argUrl := ctx.Args().Get(0)
			path := ctx.Args().Get(1)
			repoName := parseUrl(argUrl)

			if path == "" {
				path = repoName
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			fullPath, err := filepath.Abs(filepath.Join(dir, path))
			if err != nil {
				return err
			}

			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Start()

			output, err := exec.Command("git", "clone", argUrl, path).CombinedOutput()
			if err != nil {
				return fmt.Errorf("%v: %s", err, output)
			}

			defer func() {
				if err := os.RemoveAll(filepath.Join(fullPath, ".git")); err != nil {
					log.Println(err)
				}
			}()
			s.Stop()
			fmt.Printf("Project success saved on \033[32m%s\033[0m", fullPath)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func parseUrl(url string) string {
	str := strings.Replace(url, ".git", "", -1)
	arr := strings.Split(str, "/")
	return arr[len(arr)-1]
}
