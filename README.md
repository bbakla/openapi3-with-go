# Oapi-codegen
> Github repo: https://github.com/deepmap/oapi-codegen

1. Install the tool: `go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest`
2. Be sure that `GOPATH/bin` is added to the Path
3. Create the directory in which you will put the generated code. This directory is `api` in the examples below.
3. Run the command below to generate the code
```go
    oapi-codegen -package=api -generate "types,server,spec" todo-app.yaml > api/todo.gen.go 
```
4. There may be some dependencies you may need as the generated code uses them
5. The code is now under `api` directory generated. You will see an interface `ServerInterface` which you will implement.

## Features
1. *Use 3th party library*: You can use several libraries to generate handlers of server. To use `gin`:
```go
    oapi-codegen -package=api  -generate gin todo-app.yaml > api/todo.gen.go 
```

2. *Improve readability*: You can generate the code in different files to increase separately:
```
     oapi-codegen -package=api  -generate gin todo-app.yaml > api/todo-server.gen.go   
```

```
    oapi-codegen -package=api  -generate types todo-app.yaml > api/todo-type.gen.go     
```

```
    oapi-codegen -package=api  -generate client todo-app.yaml > api/todo-client.gen.go
```

```
    oapi-codegen -package=api  -generate spec todo-app.yaml > api/todo-specs.gen.go  
```
# Using config file
It could be difficult to follow to command entering parameters in the command line. The configuration files can helps us with that
1. Define a config file such as below. This is just an example. You can try all the possibilities of the configuration.
```yaml
  package: api
  output: api/server.gen.go
  generate:
    embedded-spec: false
    strict-server: true
    models: false
    chi-server: true # compatible with net/http
```
2. Run the command below:
```
    oapi-codegen -package=api --config=server-cfg.yaml  todo-app.yaml
```

We now generated the code for server side. We can also generate the models. I created the config file also for it. Feel 
free to try it

# Using openapi-generator
We could also use OpenAPI Generator's CLI tool. You can see how to install it [here](https://openapi-generator.tech/docs/installation)
After installing,  you run the command. The code will be generated in `open-api` directory. 
```
openapi-generator generate -g go-server -o open-api -i todo-app.yaml
```
You can also generate the code compatible with `echo` or `gin`. You can find more information [here](https://openapi-generator.tech/docs/generators#server-generators)
Lets generate it for `gin`:
```
openapi-generator generate -g go-gin-server -o open-api -i todo-app.yaml
```

It is also possible to define a config file not to simplify the command prompt. The list of command options can be found [here](https://openapi-generator.tech/docs/generators/go-gin-server/)

# Implementing the endpoints