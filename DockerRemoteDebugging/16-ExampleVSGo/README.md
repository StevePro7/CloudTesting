VS Code example
11-0May-2023

repeat 02-ExampleVS

01.
`Launch VS Code`

right click dockerfile
Build image
docker build --pull --rm -f "Dockerfile" -t 16examplevsgo:latest "."
right click image
Run interactive
docker run --rm -it -p 40000:40000/tcp -p 80:80/tcp 16examplevsgo:latest

Click Play extension icon
Click Play button Docker

set breakpoint
curl localhost:80
break


02.
Launch Goland IDE
change value
Build image
right click image
Run interactive

Edit configurations
Go Remote
Port 4000
Click debug button

set breakpoint
curl localhost:80
break


docker build --pull --rm -f "Dockerfile" -t 16examplevsgo:latest "."
docker run --rm -it -p 40000:40000/tcp -p 80:80/tcp 16examplevsgo:latest