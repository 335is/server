# Server
Web server that demonstrates REST API, configuration, logging, metrics, middleware, and panic handling.
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

## Run Server

```bash
./server http_address=http://localhost http_port=1111 http_content=content
```

## Routes
- [bands](http://localhost/bands)
- [band names](http://localhost/bands/names)
- [band information](http://localhost/bands/{band})
- [band members](http://localhost/bands/{band}/members)
- [band member information](http://localhost/bands/{band}/members/{name})
- [band member name](http://localhost/bands/{band}/members/{name}/name)
- [band member instruments](http://localhost/bands/{band}/members/{name}/instruments)
- [band member sings vocals](http://localhost/bands/{band}/members/{name}/instruments/vocals)
- [band member is founding member](http://localhost/bands/{band}/members/{name}/founder)
- [band members is current member](http://localhost/bands/{band}/members/{name}/current)
- [band name](http://localhost/bands/{band}/name)
- [band year started](http://localhost/bands/{band}/year)
- [metrics](http://localhost/metrics)

