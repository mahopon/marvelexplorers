#!/usr/bin/env bash

docker stop marvelfrontend
docker rm marvelfrontend
docker rmi marvelfrontend

docker build -t marvelfrontend .
docker create --name marvelfrontend -p 3000:3000 marvelfrontend
docker start marvelfrontend

echo "Container created and started up."