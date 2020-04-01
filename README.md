# graft (worker服务)

### 依赖
* go (>= 1.13.1)

```bash
export GOPATH={你的golang代码路径 /Code/go}
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

### IDE 配置

golang 需要升级到最新版
![](./doc/ide.png)

参考 ./doc/ide.png

`verdoring mode` 开启需要执行 `go mod vendor` 把依赖安装到项目目录

### 安装启动

```bash
# 修改配置文件
cp conf/app.json.example conf/app.json

./build.sh run
```

### 版本管理
`go mod`

