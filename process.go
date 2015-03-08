package main;

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"io"	
)


func main() {
	
	cmdPath := "C:\\Users\\Ajay\\Desktop\\repos\\Fitbit.NET_\\SampleDesktop\\bin\\Debug\\SampleDesktop.exe";

	simpleExecution(cmdPath);

	captureProcessOutput(cmdPath);

}

func simpleExecution(cmdPath string) {
    fmt.Printf("simpleExecution: Running " + cmdPath + "\n")
    cmd := exec.Command(cmdPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}


func captureProcessOutput(cmdPath string) { 

    fmt.Printf("captureProcessOutput: Running " + cmdPath + "\n")
	
    cmd := exec.Command(cmdPath)

	// capture the output and error pipes
	stdout, err := cmd.StdoutPipe()
	handleError(err)

	stderr, err := cmd.StderrPipe()
	handleError(err)

	err = cmd.Start()
	handleError(err)

	// Don't let main() exit before our command has finished running
	defer cmd.Wait() // Doesn't block

	// Non-blockingly echo command output to terminal
	//go io.Copy(os.Stdout, stdout) //  <---- commented out because we will print out with buff.Scan()
	go io.Copy(os.Stderr, stderr)

	// here is where we want to know the child output	

	buff := bufio.NewScanner(stdout)
	//var allText []string

	for buff.Scan() {
	     //allText = append(allText, buff.Text()+"\n")
	     fmt.Print("Output from Process: " + buff.Text() + "\n");
	}        

} 

func handleError(err error) {
	if err != nil {
	     fmt.Println(err)
	     os.Exit(1)
	}
}