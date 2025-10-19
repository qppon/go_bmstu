package main

import (
	"log"
	"os/exec"
	"runtime"
)

func main() {
	os := runtime.GOOS
	url := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	var cmd *exec.Cmd
	if os == "darwin" {
		cmd = exec.Command("open", url)
	} else if os == "windows" {
		cmd = exec.Command("cmd", "/c", "start", url)
	} else {
		cmd = exec.Command("xdg-open", url)
	}
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
