FROM scratch

MAINTAINER Eddy Hernández

WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["/app/chat", "-addr=:8080"]