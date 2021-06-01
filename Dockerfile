## 引入最新的golan ，不设置版本即为最新版本
FROM golang
## 在docker的根目录下创建相应的使用目录
RUN mkdir -p /www/webapp
## 设置工作目录
WORKDIR /www/webapp
## 把当前（宿主机上）目录下的文件都复制到docker上刚创建的目录下
COPY . /www/webapp
## 编译
# RUN go build .
# RUN go build /www/webapp
RUN go mod init main
RUN go env -w GOPROXY=https://goproxy.cn,direct

# RUN go run main.go
## 设置docker要开发的哪个端口
# EXPOSE 8080
## 启动docker需要执行的文件
# CMD go run main.go
