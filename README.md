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