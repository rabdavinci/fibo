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
3. Move cache to Redis or Memcache
4. Dockerize

## SCREENSHOTS

![Screenshot from 2022-01-11 00-19-53](https://user-images.githubusercontent.com/30826165/148826190-72f0eb17-fafa-4401-8e92-4f50ba0d5f46.png)
![Screenshot from 2022-01-11 00-21-34](https://user-images.githubusercontent.com/30826165/148826199-0b7ec61a-d8f8-40ce-9a2c-816ac5a2b50f.png)

