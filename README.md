# illuminatingdeposists-kafka


This is a sidecar project that is in progress and 
will be used in future with [illuminatingdeposists-grpc]( https://github.com/rsachdeva/illuminatingdeposits-grpc )
Currently at very early stage.

# Docker Compose Deployment

### Start Kafka with zookeeper
```shell
export COMPOSE_IGNORE_ORPHANS=True && \
docker-compose -f ./deploy/compose/docker-compose.kafka.yml up 
```

### Run producer
go run write/write.go deposit_calculations

### Run consumer
go run read/read.go deposit_calculations

### list topics
```shell
go run list/list.go
```

### Shutdown
```shell
export COMPOSE_IGNORE_ORPHANS=True && \
docker-compose -f ./deploy/compose/docker-compose.kafka.yml down
```

# Version
v0.5