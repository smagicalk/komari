FROM --platform=linux/amd64 alpine:3

WORKDIR /app

# 将 entrypoint.sh 脚本复制到容器中
COPY entrypoint.sh .
COPY komari .
# 合并命令以减少层数，并赋予执行权限
RUN apk update && \
    apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && apk del tzdata &&\
    chmod +x komari && \
    chmod +x entrypoint.sh

# 声明默认的环境变量值
# entrypoint.sh 会检查这些变量是否存在
ENV DB_TYPE=postgres
ENV BIND_ADDR=0.0.0.0:25774

EXPOSE 25774

# 【核心修正】将容器的入口点设置为您的检查脚本
ENTRYPOINT ["./entrypoint.sh"]