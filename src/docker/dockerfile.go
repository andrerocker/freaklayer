package docker

import (
    "io"
    "../util"
)

func MakeDockerBuilderImage(output io.Writer, project string, requestId string) error {
    sign := util.BuildWorkspaceSign(project, requestId)
    repository := util.ResolveWorkspaceRepositoryPath(sign)
    return util.ExecuteAndWriteToStreamFromDirectory(output, repository, "docker", "build", "-t", sign, ".")
}

func BuildPackageOnBuilderImage(output io.Writer, project string, requestId string) error {
    sign := util.BuildWorkspaceSign(project, requestId)
    repository := util.ResolveWorkspaceRepositoryPath(sign)
    args := []string { "run", "-v", "/home/andrerocker/andre/work/locaweb/infradev/docker/freaklayer/buildpacks:/buildpacks", "-v", repository+":/repository", "-t", sign, "/buildpacks/build", "/repository" }
    return util.ExecuteAndWriteToStreamFromDirectory(output, repository, "docker", args...)
}
