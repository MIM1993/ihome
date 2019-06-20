#依赖镜像
FROM alpine:3.2
ADD html /html
ADD conf /conf
ADD IhomeWeb /IhomeWeb-web

#到根目录
WORKDIR /

#执行命令
ENTRYPOINT [ "/IhomeWeb-web" ]

#端口
EXPOSE  23333