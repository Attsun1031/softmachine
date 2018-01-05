# How to generate jobapi.bp.go
```
$ protoc --version
libprotoc 3.5.1
$ protoc -I proto/ proto/jobapi.proto --go_out=plugins=grpc:.
```