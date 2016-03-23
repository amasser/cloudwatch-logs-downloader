package util

import "fmt"
import "os"
import "os/exec"
import "strings"

func PrintCommand(cmd *exec.Cmd){
        fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func PrintError(err error) {
        if err != nil{
        		fmt.Println("\n\n::::::::::::::::::Something went wrong:::::::::::::::::::")
                os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
          		fmt.Println("Something doesn't seem right, please make sure that AWS Cli is configured properly with named profile. \nFor more details: https://aws.amazon.com/cli/")
          		fmt.Println(":::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::\n\n")
        }
}

func PrintOutput(outs []byte){
        if len(outs) > 0{
                fmt.Printf("==> Output: %s\n", string(outs))
        }      
}