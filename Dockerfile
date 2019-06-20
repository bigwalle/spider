#源镜像
FROM golang

#作者
MAINTAINER Razil "292545811@qq.com"

#设置工作目录
WORKDIR $GOPATH/src/spider

COPY . .

WORKDIR cmd

RUN go build -o spider .

EXPOSE 8081

ENTRYPOINT ["./spider"]

#CMD ["/bin/bash", "build.sh"]
