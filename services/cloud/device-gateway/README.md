# Device Gateway

Rest API gateway that is used by the nodes to communicate with the cloud. 
It's a single point of contact for the nodes. 

## Concept 
Device gateway is based on code generated by [Grpc gateway](https://github.com/grpc-ecosystem/grpc-gateway). 
Microservices that are expose their interface through the gateway should [annotate their proto](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/) files with `google.api.http` annotations.

## Contribute
### Adding New Services
- Extend `gen` target of `Makefile` to generate gateway and Open API documentation for your proto file. 
You will need to change the Grpc Client import path in generated code using `sed` command. 

- run `make gen`
- add configuration section for the endpoint of grpc service in `config.go`
- add handlers of your service in main.go like this:
    ```
    err := gen.RegisterImsiServiceHandlerFromEndpoint(ctx, grcpMux, svcConf.Services.Hss, opts) 
    ```
- build and run the gateway
- check `/swagger` endpoint in browser to make sure documentation of your service is added. 

:warning: Don't forget to regenerate grpc gateway after changing service annotations in proto files.