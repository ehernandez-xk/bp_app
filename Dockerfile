FROM alpine:latest

MAINTAINER Eddy Hernández

WORKDIR /app

COPY app .

EXPOSE 8080

CMD ./app