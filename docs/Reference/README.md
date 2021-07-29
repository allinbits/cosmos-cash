# Reference Documentation

How to use the Cosmos Cash Reference Documentation.

[TOC]



## Context

In Cosmos Cash we use the  *Grand Unified Theory of Documentation* (David Laing) as described by [Divio](https://documentation.divio.com/) as a basis for our documentation strategy.

This approach outlines four specific use cases for documentation:

* [Tutorials](../Tutorials/README.md)
* [How-Tos](../How-To/README.md)
* [Explanation](./Explanation/README.md)
* [Reference](./README.md)

This section addresses the last of these: **Reference documentation**. 

For further background please see [the ADR relating to the documentation structure](./ADR/adr-002-docs-structure.md). 

## Approach

As outlined in the strategy [Reference Documentation](https://documentation.divio.com/reference/) should be **information oriented**. It should allow the reader to easily navigate the content and use in conjunction with the code.

This documentation describes the machinery, for example, classes, functions, interfaces, parameters, and so on.


## Contributing

* The content should be dry, clear and terse in style.
* All documentation should be written following [Google Documentation Best Practice](https://google.github.io/styleguide/docguide/best_practices.html)
* Generate as much documentation as possible from the code.
* Raise a PR for all documentation changes

## Layout

Reference Documentation will come in various forms:

* [Architecture diagrams](./architecture) - this will be in SVG format so that it can be stored in GitHub and version controlled. 
* [Module specifications and designs](./MODULES.md) - By convention in Cosmos SDK module documentation is stored with the module itself. However we propose that easiest way for a new user to find documentation is at the root docs folder. To accomodate this module documentation will follow existing convention, but it will be reference from this section.
* **Code-level documentation** - this will be text that is with the code and auto-generated from the code
* **API reference** - this includes REST and gRPC enpoints.
* [Glossary](,.GLOSSARY.md) - a dictionary of domain relevant terms. This can be used in conjunction with the [Cosmos Network Glossary](https://v1.cosmos.network/glossary).


## Reference

- [Google Style Guide for Markdown](https://github.com/google/styleguide/blob/gh-pages/docguide/style.md)
- [Write the Docs global community](https://www.writethedocs.org/)
- [Write the Docs Code of Conduct](https://www.writethedocs.org/code-of-conduct/#the-principles)