package main

import (
    "os"
    "./util"
    "./docker"
    "./workspace"
)

func main() {
   // sign,_ := workspace.BuildWorkspace(os.Stdout, "cegonha", "git@github.com:andrerocker/murder.git")
    _, err := workspace.BuildWorkspace(os.Stdout, "cegonha", "127.0.0.1", "git@github.com:andrerocker/rails-example.git")
    if err != nil {
        util.Message(err, os.Stdout, "finish execution")
        return
    }

    err = docker.MakeDockerBuilderImage(os.Stdout, "cegonha", "127.0.0.1")
    if err != nil {
        util.Message(err, os.Stdout, "finish execution")
        return
    }

    err = docker.BuildPackageOnBuilderImage(os.Stdout, "cegonha", "127.0.0.1")
    if err != nil {
        util.Message(err, os.Stdout, "finish execution")
        return
    }
}
