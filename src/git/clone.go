package git

import (
    "io"
    "../util"
)

func CloneProject(output io.Writer, repository string, target string) error {
    return util.ExecuteAndWriteToStream(output, "git", "clone", repository, target)
}
