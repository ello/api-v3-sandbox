FROM golang

ADD . /go/src/github.com/ello/sandbox/api-go

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgutz/dat.v1
RUN go get gopkg.in/mgutz/dat.v1/sqlx-runner
RUN go install github.com/ello/sandbox/api-go
