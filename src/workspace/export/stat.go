package export

import (
    "io"
    "fmt"
    "../util"
)

func FilesCount(output io.Writer, directory string) error {
  return util.ExecuteAndWriteToStream(output, "bash", "-c", fmt.Sprintf("find %s -type f | wc -l", directory))
}
