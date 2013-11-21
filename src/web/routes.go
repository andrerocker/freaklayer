package web

import (
    "fmt"
    "./util"
    "net/http"
    "../workspace"
    workspaceUtil "../workspace/util"
    "github.com/hoisie/web"
)

func DrawRoutes() error {
    web.Post("/workspace/(.*)/branch/(.*)", buildWorkspace)
    web.Post("/build/(.*)/image/(.*)", buildPackage)
    web.Post("/export/(.*)", buildExport)
    web.Get("/download/(.*)", download)
    return nil
}

func Run() {
    web.Run("0.0.0.0:8080")
}

func buildWorkspace(context *web.Context, project string, branch string) {
    repository := context.Params["repository"]
    requestId, writer := util.RequestIdAndResponseWriter(project, context)
    workspace.BuildWorkspace(writer, requestId, repository, branch)
}

func buildPackage(context *web.Context, project string, image string) {
    requestId, writer := util.RequestIdAndResponseWriter(project, context)
    workspace.BuildPackage(writer, requestId, image)
}

func buildExport(context *web.Context, project string) {
    requestId, writer := util.RequestIdAndResponseWriter(project, context)
    workspace.BuildExport(writer, requestId)
}

func download(context *web.Context, project string) {
    context.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=%s.tar.gz", project), true)
    requestId, _ := util.RequestIdAndResponseWriter(project, context)
    http.ServeFile(context.ResponseWriter, context.Request, workspaceUtil.ResolveWorkspaceExportFileName(requestId))
}
