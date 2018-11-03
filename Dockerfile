FROM golang:1.11
COPY . /go/src/github.com/RedVentures22/campus-2018-go
RUN go test -v github.com/RedVentures22/campus-2018-go/...
RUN go install github.com/RedVentures22/campus-2018-go/cmd/server
