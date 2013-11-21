package workspace

import (
    "io"
    "./git"
    "./util"
    "./docker"
    "./export"
)

func BuildWorkspace(output io.Writer, requestId string, repositoryUrl string, branch string) (string, error) {
    repository, _ := util.ResolveWorkspaceRepositoryPath(requestId)

    if err := util.BuildWorkspaceInitialDirs(requestId); err != nil {
        util.Message(err, output, "cannot build initial workspace structure")
        return "", err
    }

    if err := git.CloneProject(output, repositoryUrl, branch, repository); err != nil {
        util.Message(err, output, "cannot clone git repository")
        return "", err
    }

    return requestId, nil
}

func BuildPackage(output io.Writer, requestId string, image string) error {
    cache, _ := util.ResolveWorkspaceCache(requestId)
    buildpacks, _ := util.ResolveWorkspaceBuildpacksPath()
    repository, _ := util.ResolveWorkspaceRepositoryPath(requestId)

    return docker.ExecuteBuild(output, image, buildpacks, repository, cache)
}

func BuildExport(output io.Writer, requestId string) error {
    repository, _ := util.ResolveWorkspaceRepositoryPath(requestId)
    exportFile := util.ResolveWorkspaceExportFileName(requestId)

    return export.BuildTar(output, repository, exportFile)
}

func BuildDockerImage(output io.Writer, image string, content string) error {
    if err := util.BuildWorkspaceDockerInitialDirs(image); err != nil {
        util.Message(err, output, "cannot build docker workspace structure")
        return err
    }

    dockerfile := util.ResolveWorkspaceDockerImageFile(image)
    dockerfilePath := util.ResolveWorkspaceDockerImagePath(image)

    if err := docker.MakeDockerfile(output, dockerfile, content); err != nil {
        util.Message(err, output, "cannot create docker file: "+ dockerfilePath)
        return err
    }

    if err := docker.MakeImage(output, image, dockerfilePath); err != nil {
        util.Message(err, output, "cannot create docker image: "+ dockerfilePath)
        return err
    }

    return nil
}
