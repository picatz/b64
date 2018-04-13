package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func pipeGiven() bool {
	f, _ := os.Stdin.Stat()
	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	} else {
		return true
	}
}

func argumentsGiven(c *cli.Context) bool {
	if len(c.Args()) > 0 {
		return true
	} else {
		return false
	}
}

func readArguments(c *cli.Context) <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		for i := 0; i < len(c.Args()); i++ {
			messages <- c.Args().Get(i)
		}
	}()
	return messages
}

func readStdin() <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Err() != nil {
				break
			}
			data := []byte(scanner.Text())
			if len(data) > 0 {
				messages <- strings.TrimSpace(string(data))
			}
		}
	}()
	return messages
}

func encode(input string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(input)))
}

func decode(input string) string {
	decodedString, _ := base64.StdEncoding.DecodeString(input)
	return string(decodedString)
}

func main() {
	app := cli.NewApp()

	app.Name = "b64"
	app.Version = "1.0.0"
	app.Usage = "base64 encoder and decoder"

	app.Commands = []cli.Command{
		{
			Name:  "encode",
			Usage: "encode the given string(s) as base64",
			Action: func(c *cli.Context) error {
				if pipeGiven() {
					for input := range readStdin() {
						fmt.Println(encode(input))
					}
				} else if argumentsGiven(c) {
					for arg := range readArguments(c) {
						fmt.Println(encode(arg))
					}
				}
				return nil
			},
		},
		{
			Name:  "decode",
			Usage: "decode the given string(s) as base64",
			Action: func(c *cli.Context) error {
				if pipeGiven() {
					for input := range readStdin() {
						fmt.Println(decode(input))
					}
				} else if argumentsGiven(c) {
					for arg := range readArguments(c) {
						fmt.Println(decode(arg))
					}
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
