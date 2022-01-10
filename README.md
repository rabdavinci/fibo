# Fibonacci REST API

This repo contains solutions of "Fibonacci slices" REST API on Golang

## How 2RUN

1. Clone project

```
git clone git@github.com:rabdavinci/fibo.git .
```

2. Run microservice

```
$ make build
$ make run
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

## TODO

1. Add GRPC server
2. Use config package, env for keeping params
3. Use cache (Redis f. e.)
4. Dockerize
