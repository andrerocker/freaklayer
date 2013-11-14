package docker

import (
    "io"
    "fmt"
    "../util"
)

func MakeBuilderImage(output io.Writer, project string, requestId string) error {
    sign := util.BuildWorkspaceSign(project, requestId)
    repository, _ := util.ResolveWorkspaceRepositoryPath(sign)
    return util.ExecuteAndWriteToStreamFromDirectory(output, repository, "docker", "build", "-t", sign, ".")
}

func BuildPackage(output io.Writer, project string, requestId string, image string) error {
    sign := util.BuildWorkspaceSign(project, requestId)
    buildpacks, _ := util.ResolveWorkspaceBuildpacksPath()
    repository, _ := util.ResolveWorkspaceRepositoryPath(sign)

    buildpacksMount := fmt.Sprintf("%s:/buildpacks", buildpacks)
    repositoryMount := fmt.Sprintf("%s:/repository", repository)

    args := []string { "run", "-v", buildpacksMount, "-v", repositoryMount, "-t", image, "/buildpacks/build", "/repository" }
    return util.ExecuteAndWriteToStreamFromDirectory(output, repository, "docker", args...)
}
