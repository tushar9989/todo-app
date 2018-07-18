FROM golang

ADD . /go/src/github.com/tushar9989/todo-app

RUN go get github.com/julienschmidt/httprouter
RUN go install github.com/tushar9989/todo-app

ENTRYPOINT /go/bin/todo-app

EXPOSE 8080