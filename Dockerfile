# 基础镜像
FROM ubuntu:20.04
# 把编译后的打包进来这个镜像，放在工作目录 /app。随便换的
COPY webook /app/webook
WORKDIR /app
# 最佳
ENTRYPOINT ["/app/webook"]

