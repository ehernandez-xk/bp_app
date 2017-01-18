FROM scratch

MAINTAINER Eddy Hernández

WORKDIR /app

COPY . .

EXPOSE 8080

WORKDIR /app/chat

CMD ["/app/chat/chat", "-addr=:8080"]