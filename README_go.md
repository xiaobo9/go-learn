# GO GUI

## github.com/lxn/walk

```bash
go get -d https://github.com/akavel/rsrc
go install https://github.com/akavel/rsrc
rsrc -manifest main.manifest -o rsrc.syso

go build -ldflags="-H windowsgui"
```

The usual default message loop includes calls to win32 API functions, which incurs a decent amount of runtime overhead coming from Go. As an alternative to this, you may compile Walk using an optional C implementation of the main message loop, by passing the walk_use_cgo build tag:

```bash
go build -tags walk_use_cgo
```

### debug

```bash
go get -d github.com/go-delve/delve/cmd/dlv
go install github.com/go-delve/delve/cmd/dlv
```

## use ssh instead of http/https

```bash
git config --global url."git@git.xiaobo9.top:".insteadOf "http://git.xiaobo9.top/"
go env -w GOPRIVATE=git.xiaobo9.top
go get -v -insecure git.xiaobo9.top/projectTest # http请求服务，https 不用 insecure
```

### project

```bash
go mod init github.com/xiaobo9/go-learn
go mod download
go install github.com/xiaobo9/go-learn
```

### import local modules

```bash
require "github.com/userName/otherModule" v0.0.0
replace "github.com/userName/otherModule" v0.0.0 => "local physical path to the otherModule"
```

### build

```bash
GOOS=linux GOARCH=amd64 go build go-learn
# 压缩体积
go build -ldflags="-s -w" -o server main.go
# 用 upx 压缩
go build -o server main.go && upx -9 server
#用 upx 进一步压缩 
go build -ldflags="-s -w" -o server main.go && upx -9 server
```

## vscode

```json
// .vscode/launch.json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}"
            // "program": "${workspaceFolder}"
            // "program": "${workspaceFolder}/ch1/server/main/"
        }
    ]
}
```

## go 接口 方法判断

```go
// if se, ok := v.Value.(SetOptioner); ok {
// 判断是否有 SetOptioner 方法的接口，然后就可以直接调用了？
// html.go L525 renderLink
// n := node.(*ast.Link) 直接获取接口实现类的属性
for i := l - 1; i >= 0; i-- {
    v := r.config.NodeRenderers[i]
    nr, _ := v.Value.(NodeRenderer)
    if se, ok := v.Value.(SetOptioner); ok {
        for oname, ovalue := range r.options {
            se.SetOption(oname, ovalue)
        }
    }
    nr.RegisterFuncs(r)
}
```
