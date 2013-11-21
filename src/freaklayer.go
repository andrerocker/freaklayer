package main

import (
    "os"
    "./web"
    "./workspace/util"
)

// curl http:://localhost:3000/images
// curl -X POST http://localhost:3000/build/cegonha/image/locaweb-ruby

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
