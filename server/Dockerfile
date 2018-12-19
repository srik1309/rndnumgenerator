FROM golang

ADD . /go/src/github.com/srik1309/rndnumgenerator/server
WORKDIR /go/src/github.com/srik1309/rndnumgenerator

RUN go get github.com/srik1309/rndnumgenerator/server

#RUN go install server

ENTRYPOINT [ "/go/bin/server" ]

EXPOSE 50052

