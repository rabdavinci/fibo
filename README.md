# Fibonacci REST API

This repo contains solutions of "Fibonacci slices" REST API on Golang

## How 2RUN

1. Clone project

```
git clone git@github.com:rabdavinci/fibo.git .
```

2. Run microservice

```
$ docker-compose up
```

3. Calculate fibonacci slice

```
POST localhost:9090
BODY {"x":5,"y":10}
```
4. Run tests
```
$ make test
```

5. Do GRPC Request

## TODO

1. Finish logging
2. Add json validation in middleware
3. Move caching at the data level(Currently it works for HTTP api only)