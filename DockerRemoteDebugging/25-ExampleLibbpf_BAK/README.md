VS Code example
20-0May-2023

repeat 02-ExampleVS

docker build --pull --rm -f "Dockerfile" -t 25examplelibbpf:latest "."
docker run --rm -it -p 40000:40000/tcp -p 8081:8081/tcp 25examplelibbpf:latest

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
docker build --pull --rm -f "Dockerfile" -t 25examplelibbpf:latest "."
docker run --rm -it -p 40000:40000/tcp -p 8081:8081/tcp 25examplelibbpf:latest

set breakpoint
curl localhost:8081
break


docker build --pull --rm -f "Dockerfile" -t 25examplelibbpf:latest "."
docker run --rm -it -p 40000:40000/tcp -p 80:80/tcp 25examplelibbpf:latest

IMPORTANT
it seems that this line in the Dockerfile prevents packages like 'foo'
from being detected and built as part of binary..??
#ENV GO111MODULE=off
