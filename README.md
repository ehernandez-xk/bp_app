# bp_app

Chat application step by step from Blueprints book. author Mat Ryer

Go to the localhost:8080/chat if you are new, it will re-direct to the /login section, choose Github service.

## Install deps

`go get github.com/ehernandez-xk/bp_app/chat`

Yo need to install **bazaar**

### Debian

`sudo apt-get install bzr`

### Mac

`brew install bazaar`

## run app

`cd chat`

`go build -o chat`

`./chat`

## build in container

`./build_run.sh`

`docker logs -f myapp`

## flags

`./chat --help`

`./chat -host localhost -port 8080 -silent`

## env

To hide your application credentials

`export GITHUB_CLIENT_ID=xxxxxxx`

`export GITHUB_CLIENT_SECRET=xxxxxxx`

## Setup application on providers

### github

Use localhost only to test locally, if you want to conect other devices you need to use and IP

`Settings -> OAuth applications -> Register a new application`

```
    Application name: bp_chat_app
    Homepage URL: http://localhost:8080
    Authorization callback URL: http://localhost:8080/auth/callback/github
```
