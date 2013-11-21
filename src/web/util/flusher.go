package util

import (
    "strings"
    "net/http"
    "../../workspace/util"
    "github.com/hoisie/web"
)

type FlushedWriter struct {
    http.ResponseWriter
}

func (self FlushedWriter) Write(message []byte) (int, error) {
    flusher, _ := self.ResponseWriter.(http.Flusher)
    wrote, err := self.ResponseWriter.Write(message)
    flusher.Flush()
    return wrote, err
}

func Flushed(writer *web.Context) FlushedWriter {
    return FlushedWriter {writer.ResponseWriter}
}
func RequestIdAndResponseWriter(project string, context *web.Context) (string, FlushedWriter) {
    remote := strings.Split(context.Request.RemoteAddr, ":")[0]
    return util.WorkspaceSign(project, remote), Flushed(context)
}
