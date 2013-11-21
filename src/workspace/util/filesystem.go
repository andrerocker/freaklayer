package util

import (
    "os"
    "path/filepath"
)

func BuildWorkspaceInitialDirs(project string) error {
    work  := resolveWorkspacePath(project)
    cache, _ := ResolveWorkspaceCache(project)
    export, _ := ResolveWorkspaceExportPath(project)

    return makeDirs(work, cache, export)
}

func BuildWorkspaceDockerInitialDirs(image string) error {
    return makeDirs(ResolveWorkspaceDockerImagePath(image))
}

func ResolveWorkspaceBuildpacksPath() (string, error) {
    return filepath.Abs(GetConfig("workspace", "buildpack"))
}

func ResolveWorkspaceRepositoryPath(project string) (string, error) {
    return filepath.Abs(resolveWorkspacePath(project) + "/repository")
}

func ResolveWorkspaceExportPath(project string) (string, error) {
    return filepath.Abs(resolveWorkspacePath(project) + "/export")
}

func ResolveWorkspaceExportFileName(project string) string {
    export, _ := ResolveWorkspaceExportPath(project)
    return export + "/build.tar.gz"
}

func ResolveWorkspaceCache(project string) (string, error) {
    return filepath.Abs(resolveWorkspacePath(project) + "/cache")
}

func ResolveWorkspaceDockerImagesPath() string {
    return GetConfig("docker", "images")
}

func ResolveWorkspaceDockerImagePath(project string) string {
    return ResolveWorkspaceDockerImagesPath() + "/" + project
}

func ResolveWorkspaceDockerImageFile(project string) string {
    return ResolveWorkspaceDockerImagePath(project) + "/Dockerfile"
}

func resolveWorkspacePath(project string) string {
    return GetConfig("workspace", "directory") + "/" + project
}

func makeDirs(dirs ... string) error {
    for _, dir := range dirs {
        if err := os.MkdirAll(dir, 0777); err != nil {
            return err
        }
    }

    return nil
}
