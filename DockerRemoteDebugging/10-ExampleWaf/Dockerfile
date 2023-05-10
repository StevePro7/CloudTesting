FROM goboring/golang:1.16.7b7

ENV GOROOT=/usr/local/go
ENV GOCGO=/usr/local/go-cgo

ENV CGO_ENABLED=1


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .


RUN go build main.go
ENTRYPOINT /app/main
EXPOSE 8081