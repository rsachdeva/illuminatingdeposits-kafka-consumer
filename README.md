# Illuminating Deposits - Kafka consumer

This is a sidecar project that is in progress (early stage) and
currently is used with [illuminatingdeposists-grpc]( https://github.com/rsachdeva/illuminatingdeposits-grpc )
for the consumer part of log based message broker.

###### All commands should be executed from the root directory (illuminatingdeposits-kafka-consumer) of the project
(Development is WIP)

<p align="center">
<img src="./logo.png" alt="Illuminating Deposits Project Logo" title="Illuminating Deposits Project Logo" />
</p>


# Docker Compose Deployment for illuminatingdeposists-grpc kafka consumer

### Start Kafka consumer for illuminatingdeposists-grpc project 
This should be done after following steps with Docker compose deployment in illuminatingdeposists-grpc.
```shell
export COMPOSE_IGNORE_ORPHANS=True && \
docker-compose -f ./deploy/compose/docker-compose.grpc.consumer.yml up --build
```
Run illuminatingdeposists-grpc producer to see messages.

### Logs of running services (in a separate terminal):
```shell
docker-compose -f ./deploy/compose/docker-compose.grpc.consumer.yml logs -f --tail 1  
``` 

### Shutdown
```shell
export COMPOSE_IGNORE_ORPHANS=True && \
docker-compose -f ./deploy/compose/docker-compose.grpc.consumer.yml down
```


# kubernetes Deployment for illuminatingdeposists-grpc kafka consumer 

### Make docker images and Push Images to Docker Hub

```shell
docker rmi rsachdeva/illuminatingdeposits.grpc.consumer:v0.5.0

docker build -t rsachdeva/illuminatingdeposits.grpc.consumer:v0.5.0 -f ./build/Dockerfile.grpc.consumer .  

docker push rsachdeva/illuminatingdeposits.grpc.consumer:v0.5.0
```

### Deploy consumer service
```shell
kubectl apply -f deploy/kubernetes/.
```
Run illuminatingdeposists-grpc producer to see messages.

### See Logs for Consumer Messages
```shell
kubectl logs -l app=grpcconsumer -f
```
### Remove all resources / Shutdown

```shell
kubectl delete -f ./deploy/kubernetes/.
```

# Version
v0.5.30