package util

import (
    "io"
    "os/exec"
)

func ExecuteAndWriteToStream(output io.Writer, program string, args... string) error {
    command := exec.Command(program, args...)
    command.Stdout = output
    command.Stderr = output
    return command.Run()
}
