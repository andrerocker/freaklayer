package main

import (
    "os"
    "./workspace"
)

func main() {
    workspace.BuildWorkspace("git@github.com:andrerocker/rails-example.git", os.Stdout)
}
