package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"os/exec"
)

var region string
var bucket string
var uri string
var compression bool
var notify bool

func main() {
	app := cli.NewApp()
	app.Name = "s3shot"
	app.Usage = "Take a screenshot and upload it to an S3 bucket"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		{
			Name:  "Maarten Zuidhoorn",
			Email: "maarten@zuidhoorn.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "region,r",
			Usage:       "use the region `REGION`",
			Destination: &region,
		},
		cli.StringFlag{
			Name:        "bucket,b",
			Usage:       "upload the files to `BUCKET`",
			Destination: &bucket,
		},
		cli.StringFlag{
			Name:        "url",
			Value:       "",
			Usage:       "the url to prepend to the filename",
			Destination: &uri,
		},
		cli.BoolFlag{
			Name:        "compress,c",
			Usage:       "compress the image before uploading",
			Destination: &compression,
		},
		cli.BoolFlag{
			Name:        "notify,n",
			Usage:       "get a notification when uploading is finished",
			Destination: &notify,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Capture the whole screen",
			Action: func(context *cli.Context) error {
				fmt.Println("")
				return nil
			},
		},
		{
			Name:    "window",
			Aliases: []string{"w"},
			Usage:   "Capture the current active window",
			Action: func(context *cli.Context) error {
				fmt.Println("Window")
				return nil
			},
		},
		{
			Name:    "selection",
			Aliases: []string{"s"},
			Usage:   "Capture a selection",
			Action: func(context *cli.Context) error {
				image, err := runCommand("maim", "-su")

				// Ignore errors, since it likely means the user cancelled manually
				if err == nil {
					err := handleUpload(image)
					check(err)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func runCommand(command string, args ...string) ([]byte, error) {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func copyToClipboard(text string) error {
	command := exec.Command("xclip", "-in", "-selection", "clipboard")
	stdin, err := command.StdinPipe()
	if err != nil {
		return err
	}

	err = command.Start()
	if err != nil {
		return err
	}

	_, err = stdin.Write([]byte(text))
	if err != nil {
		return err
	}

	err = stdin.Close()
	if err != nil {
		return err
	}

	return command.Wait()
}

func sendNotification(location string) error {
	_, err := runCommand("notify-send", "-u low", "Screenshot uploaded", location)
	return err
}

func handleUpload(contents []byte) error {
	var image = contents

	if compression {
		compressed, err := compress(image)
		if err != nil {
			return err
		}

		image = compressed
	}

	filename := hashFile(image) + ".png"
	output, err := uploadFile(region, bucket, filename, image)
	if err != nil {
		return err
	}

	var location = output.Location
	if uri != "" {
		location = uri + filename
	}

	err = copyToClipboard(location)
	if err != nil {
		return err
	}

	if notify {
		err := sendNotification(location)
		if err != nil {
			return err
		}
	}

	return nil
}
