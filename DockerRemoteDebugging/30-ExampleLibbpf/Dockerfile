# don't use :latest for production
FROM golang:latest

RUN apt-get -y update &&  \
    apt-get -y upgrade && \
    apt-get install --no-install-recommends -y \
        git \
        ca-certificates \
        apt-utils \
        autoconf \
        automake \
        build-essential \
        git \
        libcurl4-openssl-dev \
        libgeoip-dev \
        liblmdb-dev \
        libpcre++-dev \
        libtool \
        libxml2-dev \
        libyajl-dev \
        pkgconf \
        wget \
        zlib1g-dev


ENV GOROOT=/usr/local/go
ENV GOCGO=/usr/local/go-cgo
ENV CGO_ENABLED=1

# Copying application code into the container
COPY . /app
WORKDIR /app

# Installing Delve
RUN CGO_ENABLED=1 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
#ENV GO111MODULE=off
RUN CGO_ENABLED=1 go build -gcflags "all=-N -l" -o main .
EXPOSE 80 40000
CMD [ "/go/bin/dlv", "--listen=:40000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/app/main" ]