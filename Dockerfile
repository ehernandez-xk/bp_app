FROM alpine:latest

MAINTAINER Eddy Hernández

WORKDIR /app

COPY . .

EXPOSE 8080

WORKDIR /app/chat

CMD ["/app/chat/chat"]