package docker

import (
    "io"
    "fmt"
    "../util"
)

// possivelmente depreciado!
// func MakeBuilderImage(output io.Writer, project string, requestId string) error {
//     sign := util.BuildWorkspaceSign(project, requestId)
//     repository, _ := util.ResolveWorkspaceRepositoryPath(sign)
//     return util.ExecuteAndWriteToStreamFromDirectory(output, repository, "docker", "build", "-t", sign, ".")
// }

func ExecuteBuild(output io.Writer, image string, buildpacks string, repository string, cache string) error {
    buildpacksMount := fmt.Sprintf("%s:/buildpacks", buildpacks)
    repositoryMount := fmt.Sprintf("%s:/repository", repository)
    cacheMount  := fmt.Sprintf("%s:/cache", cache)

    args := []string { "run", "-v", buildpacksMount, "-v", repositoryMount, "-v", cacheMount, "-u", "builder", "-t", image, "/buildpacks/build", "/repository", "/cache" }
    return util.ExecuteAndWriteToStreamFromDirectory(output, repository, "docker", args...)
}
