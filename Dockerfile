FROM golang:1.21

WORKDIR /app

COPY ./build/news ./

COPY ./config/dev-config.yaml /etc/task-news/

EXPOSE 5000

ENTRYPOINT ["./news"]