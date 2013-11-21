package docker

import (
    "io"
    "../util"
    "io/ioutil"
)

func MakeDockerfile(output io.Writer, file string, content string) error {
    return ioutil.WriteFile(file, []byte(content), 0644)
}

func MakeImage(output io.Writer, name, path string) error {
    return util.ExecuteAndWriteToStreamFromDirectory(output, path, "docker", "build", "-t", name, ".")
}
