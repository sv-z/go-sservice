FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN cd ./source/ && go mod download

RUN go get github.com/githubnemo/CompileDaemon

RUN cd ./source/ && go build main.go

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main