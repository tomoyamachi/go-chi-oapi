# go-chi-oapi

rest api with https://github.com/deepmap/oapi-codegen.

## directory

```shell
├── cmd
└── pkg
    ├── api : each endpoint function
    └── gen : autogen with "make gen-server"
```

## development

`make gen-server` renew a `pkg/gen` directory.

`go run cmd/main.go` runs a resty server.