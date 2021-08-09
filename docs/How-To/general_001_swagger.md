# Swagger

This how-tos are dedicated to the swagger documentation bundled with the project.

:warning: the Swagger API are provided as an handy method to query the note data but they cannot be used to submit transactions to the node.

## Generate the Swagger UI  

The Swagger UI is generated from the `protobuf` files in the `proto` folder in the root of the project

To generate or refresh the swagger UI run the following command from the project's root:

```
./scripts/protoc-swagger-gen.sh
```

Bundle the generate swagger UI as go package 

```
cd ./docs/Reference/swagger
statik -f -src=./swagger-ui
```


## Enable Swagger UI

Swagger UI is disabled by default. To enable Swagger support edit the `app.toml` file, 
stored by default in the current user home directory:

```bash
~/.cosmos-cash/config/app.toml
```

got to the *API Configuration* section and make sure the variables are set to `true` like in the example below

```toml
###############################################################################
###                           API Configuration                             ###
###############################################################################

[api]

# Enable defines if the API server should be enabled.
enable = true       #  <- this must be set to true

# Swagger defines if swagger documentation should automatically be registered.
swagger = true      #  <- this must be set to true
```

Finally restart the node.

The swagger UI should be available at the address `http://localhost:1317/swagger/`