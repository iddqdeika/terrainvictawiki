docker context use remote-engine
docker build -t imagetag .
docker stop cont
docker rm cont
docker run --name cont -p 80:8080 -d imagetag

docker context create remote ‐‐docker “host=ssh://root@95.163.242.191”
docker context use remote
docker compose build
docker compose up -d
