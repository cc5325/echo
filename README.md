# Echo

Simple golang echo service.

## How to build

```bash
go mod tidy
go build
```

## Environment variables:

* `ECHO_ADDRESS`: Server bind addres, in format `ip:port`. If none is provided, it uses `0.0.0.0:5325`