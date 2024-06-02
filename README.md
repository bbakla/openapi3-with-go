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
### Use 3th party frameworks
You can use several libraries to generate handlers of server. To use `gin`:
```go
    oapi-codegen -package=api  -generate gin todo-app.yaml > api/todo.gen.go 
```
### Generate code separately

You can generate the code in different files separately to increase readability:
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
### Using config file
It could be difficult to follow to command entering parameters in the command line. The configuration files can helps us with that
1. Define a config file such as below. This is just an example. You can try all the possibilities of the configuration.
```yaml
  package: oapi-codegen
  output: oapi-codegen/server.gen.go
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

# openapi-generator
We could also use OpenAPI Generator's CLI tool. You can see how to install it [here](https://openapi-generator.tech/docs/installation)
After installing,  you run the command. The code will be generated in `open-api` directory. You see that the generator will 
create each model and the api separately. It is quite easy to follow.
```
openapi-generator generate -g go-server -o open-api -i todo-app.yaml
```
## Features
### Use 3th party frameworks
You can also generate the code compatible with `echo` or `gin`. You can find more information [here](https://openapi-generator.tech/docs/generators#server-generators)
Lets generate it for `gin`:
```
openapi-generator generate -g go-gin-server -o open-api -i todo-app.yaml
```
### Use config file
It is also possible to define a config file to simplify the command prompt. The list of command options for gin server can be found [here](https://openapi-generator.tech/docs/generators/go-gin-server/)
The  command to execute the generation with config file is. You can also see the command in `Makefile`
```
openapi-generator generate -c openapi-generator-cfg.yaml -i todo-app.yaml             
```
There are some catches with `openapi-generator` though:
1. Config options differ based on the server generator. Check [here](https://openapi-generator.tech/docs/generators#server-generators) for 
the config options of your server generator 
2. Generated `main.go` and `go.mod` files override your original `main.go` and `go.mod` files, unless you define the
output directory with `-o` parameter. It is also not possible to change it through config options. It can be frustrating so be careful.
I tried to name all the possible package or directory names separately so that you can understand the configuration better.

# Implementing the endpoints
So we have generated the code! How will we use them with our code than?
## Oapi-codegen

## openapi-generator