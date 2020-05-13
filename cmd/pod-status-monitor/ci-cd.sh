#!/bin/sh
#build binnary
GIT_HEAD=`git rev-parse --short HEAD`
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pod-status-monitor

# build docker image
docker build .  --no-cache  --build-arg env=dev  -t  registry.intra.weibo.com/weibo_rd_algorithmplatform/pod-monitor:dev-${GIT_HEAD}
docker push  registry.intra.weibo.com/weibo_rd_algorithmplatform/pod-monitor:dev-${GIT_HEAD}

#deploy to kubernetes
kubectl set image  deployment/pod-status-monitor pod-status-monitor=registry.intra.weibo.com/weibo_rd_algorithmplatform/pod-monitor:dev-${GIT_HEAD} -n argo