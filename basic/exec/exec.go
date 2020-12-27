package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("dir")

	// if runtime.GOOS == "linux" {
	// 	cmd = exec.Command("ls")
	// }
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	e := cmd.Run()
	checkError(e)

}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
