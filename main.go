package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("arg is empty. specify a target directory.")
	}

	targetDir := args[1]
	dir, err := filepath.Abs(targetDir)
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(dir)

	initial := true
	for _, f := range files {
		name := f.Name()
		ext := filepath.Ext(name)
		if ext != ".wav" {
			continue
		}
		if initial {
			initial = false
			err := os.Mkdir(getDestination(dir), 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
		convert(dir, name)
	}
}

func getDestination(dir string) string {
	return filepath.Join(".", filepath.Base(dir))
}

func convert(dir, name string) {
	oldFile := filepath.Join(dir, name)
	ext := filepath.Ext(name)
	newFile := filepath.Join(getDestination(dir), strings.TrimSuffix(name, ext)+".mp3")
	cmd := exec.Command("ffmpeg", "-i", oldFile, "-ar", "44100", "-ab", "256k", "-acodec", "libmp3lame", "-f", "mp3", newFile)
	fmt.Println(cmd.String())
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
