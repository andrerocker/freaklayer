package util

import (
    "os"
    "path/filepath"
)

func BuildWorkspaceInitialDirs(project string) error {
    err := os.MkdirAll(resolveWorkspacePath(project), 0777)
    cache, _ := ResolveWorkspaceRepositoryCache(project)
    os.MkdirAll(cache, 0755)
    return err
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

func ResolveWorkspaceRepositoryCache(project string) (string, error) {
    return filepath.Abs(resolveWorkspacePath(project) + "/cache")
}

func resolveWorkspacePath(project string) string {
    return GetConfig("workspace", "directory") + "/" + project
}
