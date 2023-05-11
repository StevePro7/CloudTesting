# don't use :latest for production
FROM golang:latest

# Copying application code into the container
COPY . /app
WORKDIR /app

# Installing Delve
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
ENV GO111MODULE=off
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o main .
EXPOSE 80 4000
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/app/main" ]