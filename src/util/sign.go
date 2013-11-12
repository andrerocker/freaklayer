package util

import (
    "io"
    "fmt"
    "crypto/md5"
)

func BuildWorkspaceSign(chunks ... string) string {
    hash := md5.New()

    for _, chunk := range chunks {
        io.WriteString(hash, chunk)
    }

    return fmt.Sprintf("%s-%x", chunks[0], hash.Sum(nil))
}
