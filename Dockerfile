FROM scratch

MAINTAINER Eddy Hernández

WORKDIR /app

COPY app .
COPY templates templates

EXPOSE 8080

CMD ["/app/app"]