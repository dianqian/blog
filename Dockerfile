# 设置依赖的go环境
FROM golang:1.9.2

# author
MAINTAINER yongxue

# 安装bee
RUN go get -u github.com/beego/bee

# 将代码拷贝到目录
ADD . /go/src/nest
# 设置工作目录
WORKDIR /go/src/nest

EXPOSE 8080

CMD ["bee", "run"]