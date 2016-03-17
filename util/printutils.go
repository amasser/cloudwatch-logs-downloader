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
                os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
        }
}

func PrintOutput(outs []byte){
        if len(outs) > 0{
                fmt.Printf("==> Output: %s\n", string(outs))
        }      
}