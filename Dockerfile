FROM golang:1.18-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/mathapp

RUN mkdir -p "$APP_HOME"
ADD src "/go/src/mathapp/"

WORKDIR "/go/src/mathapp"
RUN go mod vendor
RUN go build -o main
EXPOSE 8099
CMD ["/go/src/mathapp/main"]