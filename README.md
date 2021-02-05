# illuminatingdeposists-kafka


This is a sidecar project that is in progress and 
will work with [illuminatingdeposists-grpc]( https://github.com/rsachdeva/illuminatingdeposits-grpc )

# Docker Compose Deployment

### Start Kafka with zookeeper
```shell
export COMPOSE_IGNORE_ORPHANS=True && \
docker-compose -f ./deploy/compose/docker-compose.kafka.yml up 
```

### Run producer
go run write/write.go deposits

### Run consumer
go run read/read.go deposits

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
v0.4