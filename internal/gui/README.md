# gui

## 环境准备

```bash
go get -d https://github.com/akavel/rsrc
go install https://github.com/akavel/rsrc
```
## 运行

```bash
cd cmd # 进入 cmd 目录
rsrc -manifest ../../../main.exe.manifest -o rsrc.syso
go build main.go
# go build -ldflags="-H windowsgui" main.go
```

```bash
# The usual default message loop includes calls to win32 API functions, which incurs a decent amount of runtime overhead coming from Go. As an alternative to this, you may compile Walk using an optional C implementation of the main message loop, by passing the walk_use_cgo build tag:
go build -tags walk_use_cgo
```
