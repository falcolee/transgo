FROM alpine:latest
#FROM scratch
ARG version 1.0
ENV TRANSGO_VERSION $version
# 在容器根目录 创建一个 目录
WORKDIR /app

VOLUME ["/app/conf","/app/cache"]

COPY ./bin/transgo-linux_amd64 /app/transgo

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 32000

# 运行golang程序的命令
ENTRYPOINT ["/app/transgo","-k","/app/conf","--api"]
