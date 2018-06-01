FROM golang

ADD . /go/src/github.com/tushar9989/hullo

RUN go get github.com/julienschmidt/httprouter
RUN go install github.com/tushar9989/hullo

ENTRYPOINT /go/bin/hullo

EXPOSE 8080