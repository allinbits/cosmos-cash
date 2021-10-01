# Reference Documentation

How to use the Cosmos Cash Reference Documentation.

- [Reference Documentation](#reference-documentation)
  - [Introduction](#introduction)
  - [Contributing](#contributing)
  - [Layout](#layout)
  - [Reference](#reference)

## Introduction

This section contains **Reference documentation** for Cosmos Cash. [Reference Documentation](https://documentation.divio.com/reference/) is intended to be **information-oriented**. Content must allow the reader to easily navigate the content and use the content in conjunction with the code.

This documentation describes the machinery, for example, classes, functions, interfaces, parameters, and so on.

For further information please see [the ADR relating to the documentation structure](./ADR/adr-002-docs-structure.md). 
## Contributing

* The content must be dry, clear, and terse in style.
* All documentation should be written following [Google Documentation Best Practice](https://google.github.io/styleguide/docguide/best_practices.html)
* Generate as much documentation as possible from the code.
* Raise a PR for all documentation changes

## Layout

Reference Documentation will come in various forms:

* [Architecture diagrams](./architecture) - Diagrams must be in SVG format so that the diagrams can remain crisp and clear at any resolution or size, stored in GitHub, and version controlled. 
* [Module specifications and designs](./MODULES.md) - By convention, module documentation in the Cosmos SDK is stored with the module itself. However, we propose that the easiest way for a new user to find documentation is to store the documentation at the root `docs` folder. To accomodate this module, the documentation will follow the existing convention, but the content will be reference from this section.
* **Code-level documentation** - The text that is part of the code and is used to auto-generate the documentation from the code.
* **API reference** - Including REST and gRPC endpoints.
* [Glossary](GLOSSARY.md) - a dictionary of domain-relevant terms. This glossary can be used in conjunction with the [Cosmos Network Glossary](https://v1.cosmos.network/glossary).

## Reference

- [Google Style Guide for Markdown](https://github.com/google/styleguide/blob/gh-pages/docguide/style.md)
- [Write the Docs global community](https://www.writethedocs.org/)
- [Write the Docs Code of Conduct](https://www.writethedocs.org/code-of-conduct/#the-principles)
