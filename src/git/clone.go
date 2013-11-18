package git

import (
    "io"
    "os"
    "../util"
)

func CloneProject(output io.Writer, repository string, branch string, target string) error {
    _, err := os.Stat(target)

    if err == nil {
        //TODO: pensar se vale apena fazer: reset --hard && pull
        os.RemoveAll(target)
    }

    return util.ExecuteAndWriteToStream(output, "git", "clone", "-b", branch, "--depth", "0", repository, target)
}
