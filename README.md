# Illuminating Deposits - gRPC

This is a sidecar project that is in progress (early stage) and
currently is used with [illuminatingdeposists-grpc]( https://github.com/rsachdeva/illuminatingdeposits-grpc )
for the consumer part of log based message broker.

###### All commands should be executed from the root directory (illuminatingdeposits-kafka-consumer) of the project
(Development is WIP)

<p align="center">
<img src="./logo.png" alt="Illuminating Deposits Project Logo" title="Illuminating Deposits Project Logo" />
</p>


# Docker Compose Deployment

### Start Kafka consumer for illuminatingdeposists-grpc project 
This should be done after following steps with Docker compose deployment in illuminatingdeposists-grpc.
```shell
export COMPOSE_IGNORE_ORPHANS=True && \
docker-compose -f ./deploy/compose/docker-compose.grpc.consumer.yml up
```

### Shutdown
```shell
export COMPOSE_IGNORE_ORPHANS=True && \
docker-compose -f ./deploy/compose/docker-compose.grpc.consumer.yml down
```

# Version
v0.5.20