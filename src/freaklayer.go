package main

import (
    "os"
    "./web"
    "./workspace/util"
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

func main() {
    if err := util.InitConfig(); err != nil {
        util.Message(err, os.Stdout, "problems loading configuration file")
        return
    }

    if err := web.DrawRoutes(); err != nil {
        util.Message(err, os.Stdout, "problems drawing routes")
        return
    }

    web.Run()
}


