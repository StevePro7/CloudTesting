Docker remote debugging example


VS Code example
05-0May-2023

https://blog.devgenius.io/remote-debugging-golang-application-in-vs-code-5215b43ebe31
OR
https://www.kenaqshal.com/blog/debugging-dockerized-go-applications


02-ExampleVS
go program
Dockerfile

docker build -t debug-app .
docker run -it -p 80:80 -p 4000:4000 --name debug-app debug-app

VScode
launch.json

Click "Run and debug" icon LHS	looks like play button
Docker

set break point
curl 127.0.0.1:80

Code breaks at breakpoint


Try again
docker ps -a
docker rm debug-app

re-run 2x commands

docker build -t debug-app .
docker run -it -p 80:80 -p 4000:4000 --name debug-app debug-app



03-ExampleCGO
Repeat Ex02 but integrate CGO
