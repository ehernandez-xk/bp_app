# bp_app

Go to the localhost:8080/chat if you are new, it will re-direct to the /login section, choose Github service.

## Install deps

`go get github.com/ehernandez-xk/bp_app/chat`

Yo need to install **bazaar**

### Debian

`sudo apt-get install bzr`

### Mac



## run app

`cd chat`

`go build -o chat`

`./chat`

## build in container

`./build_run.sh`

`docker logs -f myapp`

## flags

`./chat --help`

`./chat -addr :8080 -silent`

## env

To hide your application credentials

`export GITHUB_CLIENT_ID=xxxxxxx`

`export GITHUB_CLIENT_SECRET=xxxxxxx`
