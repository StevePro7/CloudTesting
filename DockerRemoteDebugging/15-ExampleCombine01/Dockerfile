FROM golang:latest

ENV GOROOT=/usr/local/go
ENV GOCGO=/usr/local/go-cgo
ENV CGO_ENABLED=1


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
#COPY . .
COPY . /app

# Installing Delve
RUN CGO_ENABLED=1 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

#ENV GO111MODULE=off
RUN CGO_ENABLED=1 go build -gcflags "all=-N -l" -o main .
#RUN CGO_ENABLED=1 go build -gcflags "all=-N -l" main.go
#RUN go build main.go
#ENTRYPOINT /app/main
EXPOSE 8081 40000