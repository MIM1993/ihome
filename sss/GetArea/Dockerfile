#依赖镜像
FROM alpine:3.2

#拷贝文件进入镜像
ADD GetArea-srv /GetArea-srv
ADD conf /conf

#启动命令
ENTRYPOINT [ "/GetArea-srv" ]
