# Server

Web server that demonstrates REST API, configuration, logging, metrics, middleware, and other cool features.

See main.go or the examples\client folder for how it works.


## Build
```go
go build -o server
```

## Configuration

The configuration can be specified using a config.yml file, environment variables, or command line arguments.

```yaml
---
http:
   address: http://localhost
   port: 2000
   content: content
```

```bash
export HTTP_ADDRESS=http://localhost
export HTTP_PORT=2000
export HTTP_CONTENT=content
```

```bash
./server http_address=http://localhost http_port=1111 http_content=content
```
