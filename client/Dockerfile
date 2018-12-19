FROM golang

ADD . /go/src/github.com/srik1309/rndnumgenerator
WORKDIR /go/src/github.com/srik1309/rndnumgenerator/client

RUN go get github.com/srik1309/rndnumgenerator/client

ENTRYPOINT [ "/go/bin/client" ]

EXPOSE 50052
