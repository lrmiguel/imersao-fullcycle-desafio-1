FROM golang:1.15

WORKDIR /go/src
ENV PATH="go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install build-essential

CMD ["tail", "-f", "/dev/null"]

CMD ["GOOS=linux", "go", "build", "main.go"]