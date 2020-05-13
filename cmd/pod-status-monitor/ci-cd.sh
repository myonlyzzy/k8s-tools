#!/bin/sh
#build binnary
GIT_HEAD:=$(shell git rev-parse --short HEAD)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pod-status-monitor

# build docker image
docker build .  --no-cache  --build-arg env=dev  -t pod-monitor:dev-${GIT_HEAD}
docker push myonlyzzy/pod-monitor:dev-${GIT_HEAD}

#deploy to kubernetes

kubectl apply -f pod-status-monitor-deploy.yaml -n argo
kubectl set image deployment pod-status-monitor