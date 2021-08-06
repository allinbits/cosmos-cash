# Technical setup 


## Software requirement

To build the project:

- [Golang](https://golang.org/dl/) v1.16 or greater 
- [`make`](https://www.gnu.org/software/make/) - to use Makefile targets
- `sha256sum` = to calculate build checksums

To run the [seed scripts](scripts/seeds/README.md)

- OS: linux or macOS
- [jq](https://stedolan.github.io/jq/) - json processor

For git pre-commit hooks

- python
- [pre-commit](https://pre-commit.com/#install) - git pre-commit hooks framework 
- [golangci-lint](https://github.com/golangci/golangci-lint) - linter for Go 


## Development environment setup (only for contributing)

For development it is highly recommended to install the configured pre-commit hooks by running from the project's root folder:

```sh
$ pre-commit install
$ pre-commit install --hook-type commit-msg
```

This operations should be run only once. 

## Building 

To build the project, run the `make build` command from the project's root folder. The output of the build will be generated in the `build` folder.

For cross-builds use the standard `GOOS` and `GOARCH` env vars. i.e. to build for windows:

```
GOOS=windows GOARCH=amd64 make build
```

> 💡 you can also use the default `go` command to build the project, check the content of the [Makefile](./Makefile) for reference

> ⚠️ on windows the `sha256sum` command is not present and therefore the `make build` command will build successfully but will fail to compute the checksum and copy the other resources to the build folder.

## Testing 

To test the project, run the `make test` command from the project's root folder. 

> 💡 you can also use the default `go` command to build the project, check the content of the [Makefile](./Makefile) for reference


## Installation 

To install the node client on your local machine, run the `make install` command from the project's root folder. 

> 💡 you can also use the default `go` command to build the project, check the content of the [Makefile](./Makefile) for reference

