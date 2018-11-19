FROM golang
COPY app /opt/html-Webpage-Golang-Decorated/app
WORKDIR "/opt/html-Webpage-Golang-Decorated/"
ENTRYPOINT "/opt/html-Webpage-Golang-Decorated/app"

