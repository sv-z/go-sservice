FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go get github.com/githubnemo/CompileDaemon

#RUN echo 'alias run_app1="/usr/bin/make -f /app/.img/go/Makefile env"' >> ~/.bashrc \
#  && . ~/.bashrc

# non-interactive shells you should create a small script and put it in your path, i.e.:
# https://stackoverflow.com/questions/36388465/how-to-set-bash-aliases-for-docker-containers-in-dockerfile
RUN printf '#!/bin/bash\n/usr/bin/make -f /app/.img/go/Makefile env' > /usr/bin/run_app && chmod +x /usr/bin/run_app

#ENTRYPOINT CompileDaemon --build="go build ./../cmd/apiserver//main.go" --command=./main
ENTRYPOINT run_app