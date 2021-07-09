# ADR 001: Documentation Structure

## Status

DRAFT

## Abstract

This ADR proposes a documentation strategy based on the *Grand Unified Theory of Documentation* (David Laing) as described by [Divio](https://documentation.divio.com/).

It outlines four specific use cases for documentation. Based on these use cases and other non-functional requirements, a structure is proposed that will address these concerns using GitHub as the Content Management System. 

In addition it also proposes 

- the use - and re-use - of document and format templates
- specific codeowners for documentation 
- comment and commit templates combined with githook checks

The outcome shall be focused, consistent, high quality documentation. 

## Context

Good documentation is important to the success of software projects.

*Writing excellent code doesn't end when your code compiles or even if your test coverage reaches 100%. It's easy to write something a computer understands, it's much harder to write something both a human and a computer understand. Your mission as a Code Health-conscious engineer is to write for humans first, computers second. Documentation is an important part of this skill.* [Google Documentation Best Practice](https://google.github.io/styleguide/docguide/best_practices.html)

The documentation use cases, as outlined by Divio are:

- Allow a new user to get started
- Show a user how to solve a specific problem
- Describe the machinery i.e. classes, functions, interfaces, parameters etc.
- Explanation/context for design, scope etc

![Documentation Quadrants](https://documentation.divio.com/_images/overview.png) 

It is key that the documentation is structured and written such that:

- Depending on use case the documentation can be found easily
- The documenation is written in an appropriate style for the use case
- Each type of documentation is written consistent style
- Documentation is scoped to a specific use cases i.e. a tutorial doesn't describe (beyond a relevant link) why the software works, it just teaches

Additional Documentation non-functional use cases include:

- It SHOULD BE as close to the code as reasonably practicable
- It SHOULD BE generated from code as much as possible
- It SHOULD USE a consistent format 
- It SHOULD BE useable from within the repository
- It COULD HAVE an automatic process that converts teh content to a website based on [Read The Docs](https://readthedocs.com/), [Gitbook](https://www.gitbook.com/) or other suitable hosting system.

## Decision

To address the use cases outlined in the context, this ADR proposes the following decisions:

- Use GitHub as primary content management [https://github.com/allinbits/cosmos-cash](https://github.com/allinbits/cosmos-cash)
- Markdown and LaTeX to deliver research publications

Given GitHub will form the content management system, we propose the following structure:

### Structure

The documentation structure shall use as much as possible a structure similar to the [Divio user cases](https://documentation.divio.com/introduction/).

|                 | Tutorials | How-to guides | Reference   | Explanation   |
|-----------------|-----------|---------------|-------------|-------------- |
| **Oriented to** | Learning  | A goal        | Information | Understanding | 
| **Must**        | Allow a newcomer to get started | Show how to solve a specific problem | Describe the machinery | Explain |
| **Takes the form of**    | A lesson | A series of steps | A dry description | A discursive explanation |
| **Analogy**     | Teaching a child to cook | Recipe in a cookery book | An encyclopedia article | A paper on culinary social history |

The specific implementation for Cosmos Cash SHOULD BE as per the following tree structure.

```
/
├── README
├── CONTRIBUTING
├── TECHNICAL-SETUP
├── CODEOWNERS
├── x/
|   ├── module_a/
|       ├── README
|       ├── docs/
|           ├── state
|           ├── state_transitions
|           ├── messages
├── docs/
    ├── README
    ├── CODEOWNERS
    ├── Explanation/
    |   ├── README
    |   ├── ADR/
    |   |   ├── README
    |   |   ├── PROCESS
    |   |   ├── adr-template
    |   |   ├── adr-{number}-{desc}
    |   ├── articles/
    |   |   ├── regulation-litepaper/
    |   |       ├── ARTICLE
    |   ├── research/
    |       ├── README
    |       ├── research_topic/
    ├── How-To/
    |   ├── HowToDoSomething/
    |   ├── HowToDoSomethingElse/
    ├── Reference/
    |   ├── README
    |   ├── GLOSSARY
    |   ├── MODULES
    |   ├── use-cases/
    |   |   ├── use-case-A
    |   |   ├── use-case-B
    |   ├── architecture/
    ├── Tutorials/
        ├── Tutorial_1/
        ├── Tutorial_2/
```

#### Root level documents

Root level of the repo SHALL HAVE the following files:

- **README.md** - use this for introduction and orientating the user. All README's SHOULD FOLLOW guidelines - see [GitHub guide](https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-on-github/about-readmes)
- **TECHNICAL-SETUP.md** - how to get started with the repo (could be a link to a tutorial or How-To)
    - Links to specific tooling set-up - development tools, linters etc
    - Dependencies such as [pre-commit](https://pre-commit.com/), 
    - Building the code
    - Running tests
- **CONTRIBUTING.md** - This details how to new users can contribute to the project. In specific,
    - Committing changes
    - Commit message formats (see below)
    - Raising PR's
    - Code of Conduct
- **CODEOWNERS** - although not part of documentation it will define the code maintainers who responsible for quality assurance on comments, PR's and issues.

#### Modules

In line with Cosmos SDK convention (TODO: needs reference) each module contains its own relevant documentation:

- **Module specifications** - these are document that outline state transitions x/module-name/docs/
- **Module-level README.md** e.g. x/module-name/README.md

This will be classed as reference documentation. It SHOULD BE descriptive, but explanatory. Explanations should be part of issues, Pull Requests and/or docs/explanation/architecture.

#### docs/

At docs level, this folder shall include the following files and folders:

- **README.md** - SHALL USE this for introduction and orientating the user, based on the content of this ADR and other materials.
- **CODEOWNERS** - This CodeOwners file details the reviewers for documentation folder. This SHALL INCLUDE the code maintainers in the root CODEOWNERS file plus a member of the Tendermint Technical Writing Team.

#### docs/Reference

Reference documentation includes a number of different forms:

- **README.md** - This document would outline the purpose of the reference documentation as per the methodology above. In addition it will also links to documentation created from the code itself, specifically
    - Code Documentation in form of Go Docs
    - Swagger API documentation
- **GLOSSARY.md** - It is key that this is reviewed regulary and applied consistently. These form the terms of reference for users and ensure that discussion and design are based on consistent terms of reference.
- **MODULES.md** - A markdown document that has references to module relevant documentation

##### docs/Reference/use-cases

This will be a folder that describe Cosmos Cash use cases. Ideally these will be in BDD format. As perscribed above, this information should be dry in nature and avoid explanation the should be cover in the explanation documentation

##### docs/Reference/architecture

This will be a folder containing architecture diagrams such as component, activity and sequence diagrams as relevant. Specifically these should be in a format suitable for version management and esay to update. Therefore these should be in SVG or DOT format and not image formats (JPEG, PNG etc.)

#### docs/Explanation

This folder will provide context for readers and are discursive in nature. Please see the [Divio Explanation page](https://documentation.divio.com/explanation/#) for more detail.

- **docs/explanation/README.md** - This will orient the reader and explain the content. 

##### docs/Explanation/ADR

This is a folder that will track decisions regrading design and architecture (such as this). It will contain the following:

- **docs/explanation/adr/README** - introduction to ADR
- **docs/explanation/adr/PROCESS.md** which details how to raise ADRs
- **docs/explanation/adr/adr-template.md** - template for raising ADR
- **docs/explanation/adr/adr-{number}-{desc}.md** - an ADR document

##### docs/Explanation/articles

This SHALL BE a folder that contains a sub-folder for each published article. By published this COULD REFER to blog posts. The folder should be name such that it describes the article purpose. Each sub-folder SHALL CONTAIN all the content relevant to the article (e.g. images, bibliographies etc). These articles can be converted into pdf using pandoc. 

In order to do this

- There SHOULD BE a makefile with targets for calling pandoc - Note: the process for building pdf are not part of the commit or release processes, but ad-hoc
- There SHOULD BE a latex template file that can create pdf files with consistent look and feel. This COULD BE the [Eisvogel template](https://github.com/Wandmalfarbe/pandoc-latex-template) with suitable modifications
- The makefile and template should be independent of the article
- THere SHOULD BE a README.md that describes how to use and build articles

> **Note:** Explanations can come in other forms, particularly issue discussion and Pull Requests.

#### docs/Tutorials

As indicated in the overview, these SHALL BE documents that target beginners and will take a user step-by-step through a process with the aim of achieving some goal. Please see the [Divio tutorial page](https://documentation.divio.com/tutorials/) for detail.

- There SHALL BE a folder for each tutorial - see the [Cosmos SDK tutorials](https://github.com/cosmos/sdk-tutorials) as an example
- The folder SHALL CONTAIN all the content relevant for that tutorial. 
- The content SHOULD BE consistent in format with [SDK tutorials](https://tutorials.cosmos.network/). 

#### docs/How-To

Compared to tutorials, [How-Tos](https://documentation.divio.com/how-to-guides/) are help an experienced reader solve a specific problem. These SHALL USE templates similar to the tutorials - see above.

### Templates

Documentation SHOULD USE templates for documentation such as module messages etc. 

- [The good docs project](https://github.com/thegooddocsproject)
- [Readme editor](https://readme.so/editor)


#### Code Comments

These are included as comments also form part of documentation. Comments SHALL FOLLOW recommendation as per [Conventional Comments](https://conventionalcomments.org/)

```
<label> [decoration]: <subject>

[discussion]
```

where `label = (praise|nitpick|suggestion|issue|question|thought|chore)`

#### Commit Comments

Commits comments will also follow a similar format as laid out following standard defined by [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/#summary). This SHOULD BE enforced as part of [pre-commit](https://pre-commit.com/) checks

## Consequences

This section describes the resulting context, after applying the decision. 

### Backwards Compatibility

Once this ADR is implemented existing documentation will be migrated from existing sources that include:

- Notion
- Other Git repos
- Published papers
- Blog posts 

### Positive

As a result of this

- Content will follow best practice that will be easy to navigate and read
- Content will be in a consistent format
- Commits, Issues and Pull Requests will follow best practice
- CHANGELOG and release documentation will benefit from better commit messages, reducing developer effort

### Negative

- There may be more effort required
- Moving modules into new repos may cause inconsistenties in the repo

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section should contain a summary of issues to be solved in future iterations (usually referencing comments from a pull-request discussion).

Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

## References

- [Google Style Guide for Markdown](https://github.com/google/styleguide/blob/gh-pages/docguide/style.md)
- [Write the Docs global community](https://www.writethedocs.org/)
- [Write the Docs Code of Conduct](https://www.writethedocs.org/code-of-conduct/#the-principles)

