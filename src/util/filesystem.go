package util

import (
    "os"
    "path/filepath"
)

func BuildWorkspaceInitialDirs(project string) error {
    return os.MkdirAll(resolveWorkspacePath(project), 0755)
}

//TODO: Possivelmente desnecessario
func CheckRepositoryDockerfile(repository string) error {
    _, err := os.Stat(repository + "/Dockerfile")
    return err
}

func ResolveWorkspaceBuildpacksPath() (string, error) {
    return filepath.Abs(GetConfig("workspace", "buildpack"))
}

func ResolveWorkspaceRepositoryPath(project string) (string, error) {
    return filepath.Abs(resolveWorkspacePath(project) + "/repository")
}

func resolveWorkspacePath(project string) string {
    return GetConfig("workspace", "directory") + "/" + project
}
