# Single

## RPC
```
docker build . -f ./rpc/Dockerfile -t shorturlrpc

docker run -d --name="shorturlrpc" -p 9001:9001 --link Etcd-server:Etcd-server --link my-mysql:my-mysql  shorturlrpc 
```
 

## kafka consumer

```
docker build . -f ./kafkaqueue/Dockerfile -t shorturlkafkaqueue

docker run -d --name="shorturlkafkaqueue" --link my-mysql:my-mysql  shorturlkafkaqueue 
```

## api
```
docker build . -f ./api/Dockerfile -t shorturlapi

docker run -d --name="shorturlapi" -p 8888:8888 --link Etcd-server:Etcd-server --link my-mysql:my-mysql  shorturlapi 
```

# docker-compose
```
docker-compose -f docker-compose.yml up -d
```

# etcd
```
docker network create etcdnetwork


docker run -d --name Etcd-server \
    --network etcdnetwork \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
```

# kafka topic
```
# 进入kafka容器
docker exec -it 0bdb6d691f86 /bin/sh

kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 2 --partitions 1 --topic short_url_micro_kafka_topic
```

# DEMO
```
POST https://t.wuguozhang.com/api/appKey/register
BODY {"appId":"123456","name":"my-test","adminSecret":"123456"}

POST https://t.wuguozhang.com//api/appKey/login
BODY {"appId":"123456","appSecret":"7FTgFPfjUCo6cI2iQvVRKiek0xVztfTLy"}

POST https://t.wuguozhang.com//api/shorten
BODY {"long":"https://www.wuguozhang.com/archives/139/"}
```