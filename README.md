# WiemanImages

This is a simple REST API server that serves to supply images to my websites, it is used to store and serve files with image optimization in mind.

It is written in [Go](https://go.dev/) using the [Gin framework](https://gin-gonic.com/docs/).

I have used [this project](https://github.com/velopert/gin-rest-api-sample) as a template to build off

## Testing

### Mocking

Mocking scripts can be generated using the `mockery` tool. It can be installed using the following docker image:
```bash
docker pull vektra/mockery
```

```shell
docker run -v ${PWD}:/src -w /src vektra/mockery --all
```


test