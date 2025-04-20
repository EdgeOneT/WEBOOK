# 定义函数来模拟 Makefile 的 docker 目标
function docker {
    # 删除 webook 文件，若文件不存在则忽略错误
    try {
        Remove-Item -Path .\ -ErrorAction Stop
    } catch {
        # 忽略错误
    }

    # 整理 Go 模块
    go mod tidy

    # 设置环境变量并构建 Go 项目
    $env:GOOS = "linux"
    $env:GOARCH = "arm"
    go build -o webook .

    # 删除 Docker 镜像
    docker rmi -f edge/webook:v0.0.1

    # 构建 Docker 镜像
    docker build -t edge/webook:v0.0.1 .
}

# 调用函数
docker