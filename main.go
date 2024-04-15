package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/mattn/go-tty"
	"github.com/urfave/cli/v2"
)

func Get_SecHackX(X int) string {
	return fmt.Sprintf("SecHack%03d", X)
}

/*
func success_animation(result int) {
	for i := 0; i < 3; i++ {
		color.New(color.BgRed, color.FgWhite).Printf("\r%s", Get_SecHackX(result))
		time.Sleep(time.Millisecond * 1000)
		color.New(color.BgWhite, color.FgRed).Printf("\r%s", Get_SecHackX(result))
		time.Sleep(time.Millisecond * 1000)
	}
}
*/

func main() {
	app := &cli.App{
		Name:  "SecHack365-Omikuzi-CLI",
		Usage: "Omikuzi-CLI of SecHack365. Confirm with e key.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "cracking-result",
				Value: -1,
				Usage: "Crack the result",
			},
		},
		Action: func(c *cli.Context) error {
			tty, err := tty.Open()
			if err != nil {
				log.Fatal(err)
			}

			defer tty.Close()

			flag := false
			go (func() {
				for {
					v, err := tty.ReadRune()
					if err != nil {
						log.Fatal(err)
					}
					if v == 'e' {
						flag = true
						return
					}
				}
			})()

			result := 0
			for !flag {
				fmt.Printf("\r%s", Get_SecHackX(result))
				result = (result + 5) % 400
				time.Sleep(time.Millisecond * 1)
			}
			if c.Int("cracking-result") >= 0 {
				result = c.Int("cracking-result")
			}
			fmt.Printf("\r%s", Get_SecHackX(result))
			fmt.Println()
			fmt.Println()

			if result == 365 {
				color.Success.Println("Success! Have nice 365 days!")
			} else {
				color.New(color.BgRed, color.FgBlack).Println("Failure!")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
