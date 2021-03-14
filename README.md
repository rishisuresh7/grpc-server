# grpc-server
## Basic GRPC server for CRUD operations

### Prerequisites
    - make
    - go

###### Config Usage(to be set at env)

- `PORT` -  Port number the application will start on
- `TOKEN` - Token for proxy authentication
- `MONGO_URI` - Host and port for mongodb service
- `MONGO_USERNAME` - Username for mongodb
- `MONGO_PASSWORD` - Password for mongodb

#### Clean binaries
```shell
$ make clean
```

#### Run all tests
```shell
$ make test
```

#### Build application
```shell
$ make build
```

#### Run application with default config
```shell
$ make run
```

#### Build application for linux
```shell
$ make build-linux
```
