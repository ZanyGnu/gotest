package main;

import (
	"fmt"
	"os"
	"os/exec"
)


func main() {
	cmdPath := "C:\\Users\\Ajay\\Desktop\\repos\\Fitbit.NET_\\SampleDesktop\\bin\\Debug\\SampleDesktop.exe";
    cmd := exec.Command(cmdPath)
    fmt.Printf("Running " + cmdPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}