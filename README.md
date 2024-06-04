## Proto Generation

**Generate Client**
```shell
protoc --go_out=./internal/webhook ./internal/webhook/api/*
```

**Generate Server**
```shell
protoc --go-grpc_out=./internal/webhook ./internal/webhook/api/* 
```