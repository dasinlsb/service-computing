set -ex

DOCKER_COMPOSE_VERSION=1.25.0

sudo rm /usr/local/bin/docker-compose
curl -L https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-`uname -s`-`uname -m` > docker-compose
chmod +x ./docker-compose
./docker-compose --version
sudo mv docker-compose /usr/local/bin

cd simple-web
pwd
sudo docker-compose pull
sudo docker-compose build
