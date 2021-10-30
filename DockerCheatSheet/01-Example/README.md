##### README.md
###### 31/08/2021
```
00. Setup
touch app.py
touch Dockerfile
```
```
01. Docker Image
docker build -t flask_app:0.1 .
docker images
```
```
02. Docker Container
docker run -d -p 5000:5000 flask_app:0.1
docker ps
```
```
03. Port in use
netstat -tulnap | grep 5000
sudo apt install net-tools
sudo fuser -k 5000/tcp
netstat -tulnap | grep 5000
```
```
04. Port in use
curl http://localhost:5000
```
```
05. Cleanup
docker stop $(docker ps -q)
docker rm $(docker ps -aq)
```
