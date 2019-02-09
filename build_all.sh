#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-msgstore .
docker build -t angelalonso/go-msgstore -f Dockerfile .
docker login
docker push angelalonso/go-msgstore:latest
cd helm
helm install --name go-msgstore --namespace app .


