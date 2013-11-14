package util

import (
    "io"
    "fmt"
    "os/exec"
)

func ExecuteAndWriteToStream(output io.Writer, program string, args... string) error {
    return ExecuteAndWriteToStreamFromDirectory(output, "", program, args...)
}

func ExecuteAndWriteToStreamFromDirectory(output io.Writer, directory string, program string, args ... string) error {
    command := exec.Command(program, args...)
    command.Dir = directory
    command.Stdout = output
    command.Stderr = output
    fmt.Println(command)
    return command.Run()
}
