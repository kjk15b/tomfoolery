package main

import (
	"fmt"
	"os/exec"
)

func main() {
	load := exec.Command("echo hello")
	remove := exec.Command("rm main")
	if err := load.Run(); err != nil {
		fmt.Println("ERROR: ", err)
	}

	if err := remove.Run(); err != nil {
		fmt.Println("ERROR: ", err)
	}
}