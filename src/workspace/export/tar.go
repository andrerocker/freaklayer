package export

import (
    "io"
    "../util"
)

func BuildTar(output io.Writer, directory string, target string) error {
  return util.ExecuteAndWriteToStreamFromDirectory(output, directory, "tar", "--exclude", ".git", "-czvf", target, ".")
}
