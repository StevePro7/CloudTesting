FROM golang:latest

ENV GOROOT=/usr/local/go
ENV GOCGO=/usr/local/go-cgo
ENV CGO_ENABLED=1


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
#COPY . .
COPY . /app


RUN go build main.go
ENTRYPOINT /app/main
EXPOSE 8081 40000