package util

import "os"

func BuildWorkspaceInitialDirs(project string) error {
    return os.MkdirAll(resolveWorkspacePath(project), 0755)
}

func CheckRepositoryDockerfile(repository string) error {
    _, err := os.Stat(repository + "/Dockerfile")
    return err
}

func ResolveWorkspaceRepositoryPath(project string) string {
    return resolveWorkspacePath(project) + "/repository"
}

func resolveWorkspacePath(project string) string {
    // read from config file
    return "/tmp/freaklayer/project/" + project
}
