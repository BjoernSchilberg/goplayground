# README

## Requirements

```shell
go get -u github.com/graphql-go/graphql
```

## Example query

```shell
curl -g 'http://localhost:3000/graphql?query={songs{id,album,title,duration}}'
```