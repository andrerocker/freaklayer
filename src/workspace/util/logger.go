package util

import "io"

func Message(err error, output io.Writer, message string) {
    io.WriteString(output, message+"\n")
}
