# Docker

## Build your Go program

```shell
go build
```

## Build your Golang Docker container image


```shell
docker image build -t hello .
```

## Run your container image

```shell
docker container run -p 8888:8888 hello
```
