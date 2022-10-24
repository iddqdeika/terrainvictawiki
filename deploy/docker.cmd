docker context use remote-engine
docker build -t imagetag .
docker stop cont
docker rm cont
docker run --name cont -p 80:8080 -d imagetag
