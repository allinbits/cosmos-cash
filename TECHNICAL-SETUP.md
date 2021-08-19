# Technical setup 


## Software requirement

To build the project:

- [Golang](https://golang.org/dl/) v1.16 or higher 
- [`make`](https://www.gnu.org/software/make/) to use Makefile targets
- `sha256sum` to calculate build checksums

To run the [seed scripts](scripts/seeds/README.md):

- OS: Linux or macOS
- [jq](https://stedolan.github.io/jq/) JSON processor

For git pre-commit hooks:

- Python v3.8 or higher
- [pre-commit](https://pre-commit.com/#install) git pre-commit hooks framework 
- [golangci-lint](https://github.com/golangci/golangci-lint) linter for Go 

### Swagger 

To generate the Swagger UI for the node REST endpoints:

- [`swagger-combine`](https://www.npmjs.com/package/swagger-combine)


## Development environment setup

Install the configured pre-commit hooks by running these commands from the project's root folder:

```sh
$ pre-commit install
$ pre-commit install --hook-type commit-msg
```

## Building 

To build the Cosmos Cash node and command line client, run the `make build` command from the project's root folder. The output of the build will be generated in the `build` folder.

For cross-builds use the standard `GOOS` and `GOARCH` env vars. i.e. to build for windows:

```
GOOS=windows GOARCH=amd64 make build
```

> 💡 You can also use the default `go` command to build the project. For reference, see the contents of the [Makefile](./Makefile).

> ⚠️ on Windows, the `sha256sum` command is not present so although the `make build` command builds successfully, it fails to compute the checksum and does not copy the other resources to the build folder.

## Testing 

To run the unit and integration tests, run the `make test` command from the project's root folder. 

> 💡 you can also use the default `go` command to build the project, check the content of the [Makefile](./Makefile) for reference


## Installation 

To install the node client on your local machine, run the `make install` command from the project's root folder. 

> 💡 you can also use the default `go` command to build the project, check the content of the [Makefile](./Makefile) for reference
