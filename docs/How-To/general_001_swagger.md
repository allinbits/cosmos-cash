# Swagger

This how-to section is dedicated to the Swagger documentation that is bundled with the project.

**Note:** The Swagger API is supported only as a useful method to query the node data. The Swagger API cannot be used to submit transactions to the node.

## Generate the Swagger UI  

The Swagger UI is generated from the `protobuf` files in the `proto` folder in the root of the project

To generate or refresh the swagger UI, run the following command from the project's root folder:

```
./scripts/protoc-swagger-gen.sh
```

Bundle the generate swagger UI as go package:

```
cd ./docs/Reference/swagger
statik -f -src=./swagger-ui
```


## Enable Swagger UI

Swagger UI is disabled by default. Use the `app.toml` file to enable Swagger support. The default location of the `app.toml` file is the current user home directory:

```bash
~/.cosmos-cash/config/app.toml
```

In the *API Configuration* section, make sure the variables are set to `true` as shown in the following example:

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
