package docker

import (
    "io"
    "fmt"
    "../util"
)


func ExecuteBuild(output io.Writer, image string, buildpacks string, repository string, cache string) error {
    buildpacksMount := fmt.Sprintf("%s:/buildpacks", buildpacks)
    repositoryMount := fmt.Sprintf("%s:/repository", repository)
    cacheMount  := fmt.Sprintf("%s:/cache", cache)

    args := []string { "run", "-v", buildpacksMount, "-v", repositoryMount, "-v", cacheMount, "-u", "builder", "-t", image, "/buildpacks/build", "/repository", "/cache" }
    return util.ExecuteAndWriteToStreamFromDirectory(output, repository, "docker", args...)
}
