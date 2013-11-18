package main

import (
    "os"
    "./util"
    "./docker"
    "./workspace"
)

// curl http:://localhost:3000/images
// curl -X POST http://localhost:3000/workspace/cegonha -d 'repo: git@bacon.com:bacon.git'
// curl -X POST http://localhost:3000/build/cegonha/image/locaweb-ruby
func main() {
    err := util.InitConfig()
    if err != nil {
        util.Message(err, os.Stdout, "problems loading configuration file")
        return
    }

    //_, err = workspace.BuildWorkspace(os.Stdout, "cegonha", "127.0.0.1", "git@github.com:andrerocker/rails-example.git")
    _, err = workspace.BuildWorkspace(os.Stdout, "cegonha", "127.0.0.1", "git@code.locaweb.com.br:iaas/cegonha.git")
    if err != nil {
        util.Message(err, os.Stdout, "finish execution")
        return
    }

    // err = docker.MakeBuilderImage(os.Stdout, "cegonha", "127.0.0.1")
    // if err != nil {
    //     util.Message(err, os.Stdout, "finish execution")
    //     return
    // }

    err = docker.BuildPackage(os.Stdout, "cegonha", "127.0.0.1", "locaweb-ruby")
    if err != nil {
        util.Message(err, os.Stdout, "finish execution")
        return
    }
}
