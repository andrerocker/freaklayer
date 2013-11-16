package workspace

import (
    "io"
    "../git"
    "../util"
)

func BuildWorkspace(output io.Writer, project string, requestId string, repositoryUrl string) (string, error) {
    sign := util.BuildWorkspaceSign(project, requestId)
    repository, _ := util.ResolveWorkspaceRepositoryPath(sign)

    if err := util.BuildWorkspaceInitialDirs(sign); err != nil {
        util.Message(err, output, "cannot build initial workspace structure")
        return "", err
    }

    if err := git.CloneProject(output, repositoryUrl, repository); err != nil {
        util.Message(err, output, "cannot clone git repository")
        return "", err
    }

    return sign, nil
}
