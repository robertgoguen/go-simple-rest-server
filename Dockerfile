FROM golang:latest

RUN apt-get update && apt-get install -y go-dep

WORKDIR /opt/workspace
RUN git clone https://github.com/robertgoguen/go-simple-rest-server.git go-simple-rest-server/go/src/github.com/robertgoguen/go-simple-rest-server

ENV GOPATH     /opt/workspace/go-simple-rest-server/go
ENV GOPATH_SRC /opt/workspace/go-simple-rest-server/go/src
ENV GOBIN      /opt/workspace/go-simple-rest-server/go/bin

WORKDIR /opt/workspace/go-simple-rest-server/go/src/github.com/robertgoguen/go-simple-rest-server

RUN dep ensure -v

RUN CGO_ENABLED=1 go install .

EXPOSE 8080

CMD /opt/workspace/go-simple-rest-server/go/bin/go-simple-rest-server
