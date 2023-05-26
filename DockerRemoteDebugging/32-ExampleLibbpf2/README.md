VS Code example
25-0May-2023

repeat 02-ExampleVS

docker build --pull --rm -f "Dockerfile" -t 32examplelibbpf:latest "."
docker run --rm -it -p 40000:40000/tcp -p 8081:8081/tcp 32examplelibbpf:latest


=> ERROR [5/5] RUN CGO_ENABLED=1 go build -gcflags "all=-N -l" -o main .                                       4.8s
------                                                                                                               
> [5/5] RUN CGO_ENABLED=1 go build -gcflags "all=-N -l" -o main .:                                                  
#10 4.717 # testwebapi                                                                                               
#10 4.717 /usr/local/go/pkg/tool/linux_amd64/link: running gcc failed: exit status 1                                 
#10 4.717 /usr/bin/ld: cannot find -lelf                                                                             
#10 4.717 /usr/bin/ld: cannot find -lz
#10 4.717 collect2: error: ld returned 1 exit status
#10 4.717
------
Dockerfile:15
--------------------
13 |     RUN CGO_ENABLED=1 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
14 |     #ENV GO111MODULE=off
15 | >>> RUN CGO_ENABLED=1 go build -gcflags "all=-N -l" -o main .
16 |     EXPOSE 80 40000
17 |     CMD [ "/go/bin/dlv", "--listen=:40000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/app/main" ]
--------------------
error: failed to solve: process "/bin/sh -c CGO_ENABLED=1 go build -gcflags \"all=-N -l\" -o main ." did not complete successfully: exit code: 1



IMPORTANT
curl localhost:8081
curl: (7) Failed to connect to localhost port 8081: Connection refused

Repeat as per 20-example
install libbpf as per notes here
https://github.com/StevePro7/FelixTesting/blob/main/Felix/anNotes.txt
i.e.
LIBBPF
cd bpf-gpl/include
git clone https://github.com/libbpf/libbpf
cd libbpf/src
make
sudo make install


01.
Launch Goland IDE
Edit configurations
+
Go Remote
Port 40000
OK
click Unnamed Debug icon

set breakpoint
curl localhost:8081
break


Launch VS Code
Click Play extension icon
Click Play button Docker


02.
Launch VS Code
docker build --pull --rm -f "Dockerfile" -t 32examplelibbpf:latest "."
docker run --rm -it -p 40000:40000/tcp -p 8081:8081/tcp 32examplelibbpf:latest

set breakpoint
curl localhost:8081
break


docker build --pull --rm -f "Dockerfile" -t 32examplelibbpf:latest "."
docker run --rm -it -p 40000:40000/tcp -p 80:80/tcp 32examplelibbpf:latest

IMPORTANT
it seems that this line in the Dockerfile prevents packages like 'foo'
from being detected and built as part of binary..??
#ENV GO111MODULE=off
