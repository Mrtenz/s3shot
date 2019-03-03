package main

import (
	"io/ioutil"
	"os/exec"
)

const UncompressedName = "/tmp/screenshot.png"
const CompressedName = "/tmp/screenshot-compressed.png"

func compress(image []byte) ([]byte, error) {
	err := ioutil.WriteFile(UncompressedName, image, 0600)
	if err != nil {
		return nil, err
	}

	command := exec.Command("pngcrush", UncompressedName, CompressedName)
	err = command.Run()
	if err != nil {
		return nil, err
	}

	compressed, err := ioutil.ReadFile(CompressedName)
	if err != nil {
		return nil, err
	}

	return compressed, nil
}
