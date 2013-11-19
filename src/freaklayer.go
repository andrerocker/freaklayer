package main

import (
    "os"
    "fmt"
    "strings"
    "net/http"
    "./workspace"
    "./workspace/util"
    "github.com/hoisie/web"
)

// curl http:://localhost:3000/images
// curl -X POST http://localhost:3000/build/cegonha/image/locaweb-ruby

// func main() {
//     if err := util.InitConfig(); err != nil {
//         util.Message(err, os.Stdout, "problems loading configuration file")
//         return
//     }
//
//     requestId := util.WorkspaceSign("cegonha", "127.0.0.1")
//     _, err := workspace.BuildWorkspace(os.Stdout, requestId, "git@code.locaweb.com.br:iaas/cegonha.git")
//     if err != nil {
//         util.Message(err, os.Stdout, "finish execution")
//         return
//     }
//
//    err = workspace.BuildPackage(os.Stdout, requestId, "locaweb-ruby")
//    if err != nil {
//        util.Message(err, os.Stdout, "finish execution")
//        return
//    }
//
//    err = workspace.BuildExport(os.Stdout, requestId)
//     if err != nil {
//         util.Message(err, os.Stdout, "finish execution")
//         return
//     }
// }

type Acme struct {
    http.ResponseWriter
}

func (self Acme) WriteString(message string) {
    self.WriteString(message)
    self.Flush()
}

func initialize() {
    if err := util.InitConfig(); err != nil {
        util.Message(err, os.Stdout, "problems loading configuration file")
        return
    }
}

func main() {
    initialize()
    web.Post("/workspace/(.*)/branch/(.*)", buildWorkspace)
    web.Post("/build/(.*)/image/(.*)", buildPackage)
    web.Run("0.0.0.0:8080")
}

func buildWorkspace(context *web.Context, project string, branch string) {
    acme := Acme{&context.ResponseWriter}
    repository := context.Params["repository"]
    requestId := resolveRequestId(project, context)
    workspace.BuildWorkspace(acme, requestId, repository, branch)
}

func buildPackage(context *web.Context, project string, image string) {
    requestId := resolveRequestId(project, context)
    err := workspace.BuildPackage(context.ResponseWriter, requestId, image)
    fmt.Println(err)
}

func resolveRequestId(project string, context *web.Context) string {
    remote := strings.Split(context.Request.RemoteAddr, ":")[0]
    return util.WorkspaceSign(project, remote)
}

 //func (ctx *web.Context) WriteString(content string) {
 //    fmt.Println("YEAHHHH")
 //    ctx.ResponseWriter.Write([]byte(content))
 //}
