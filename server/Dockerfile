FROM golang

ADD . /go/src/github.com/srik1309/rndnumgenerator
WORKDIR /go/src/github.com/srik1309/rndnumgenerator/server

RUN go get github.com/srik1309/rndnumgenerator/server

ENTRYPOINT [ "/go/bin/server" ]

EXPOSE 50052