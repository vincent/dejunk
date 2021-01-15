FROM golang:latest
ADD . /go/src/github.com/vincent/dejunk
WORKDIR /go/src/github.com/vincent/dejunk
RUN go install -v
ENTRYPOINT ["dejunk"]
