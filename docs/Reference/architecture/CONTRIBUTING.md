# CONTRIBUTING

## Installation

Cosmos Cash diagrams are based on PlantUML. 

PlantUML can be installed on your local system through your platforms' package manager:

MacOS, through [Homebrew](https://brew.sh/)

```
> brew install plantuml
```

Windows, through [Chocolatey](https://chocolatey.org/)

```
> choco install plantuml
```

Linux-based systems, through your distributions package manager

## Contributing diagrams

The source code for all diagrams SHOULD be stored in `./src` folder AND have the `.puml` suffix. Please refer to the [PlanUML documentation](https://plantuml.com) for further details.

## Themes

Where possible all diagrams will be rendered to have a similar look and feel. To achieve this add the following to your diagram:

```
@startuml myDiagram
!theme tendermint from ../themes/
...

@enduml
```

## How to generate diagram images

There are [many ways to run PlantUML](https://plantuml.com/running) to generate images from the `.puml` files, but this project prefers the following: 

* PlantUML CLI
* Embed PlantUML diagrams using PlantUML's proxy service in Markdown
* Visual Studio Code Plugin - best for development

### PlantUML CLI

There is [makefile](.makefile) that will build all `src/*.puml` files. To create SVG images just run `make (svg|png)` as needed. `make clean` will get a clean folder ready for new build.

This is the preferred and cleanest way to create all the architecture documentation.

### Proxy service

This method uses PlantUML's [proxy service](https://plantuml.com/server) to put generate images and embed into a Markdown page. 

To use the proxy service integration simply use:

```
![cached image](http://www.plantuml.com/plantuml/proxy?src=https://raw.github.com/plantuml/plantuml-server/master/src/main/webapp/resource/test2diagrams.txt)
```
Or if caching is required:
```
![uncached image](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.github.com/plantuml/plantuml-server/master/src/main/webapp/resource/test2diagrams.txt)
```

### Visual Studio Code

If you use [Visual Studio Code](https://code.visualstudio.com/) there is [PlantUML plugin](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml) available. 

From here, one can preview diagrams through the command palette (CTRL+SHIFT+P on Windows/Linux or COMMAND+SHIFT+P on MacOS)

```
> PlantUML: Preview Current Diagram
```



## Reference

* [Markdown native diagrams with PlantUML](https://blog.anoff.io/2018-07-31-diagrams-with-plantuml/)
* [PlantUML Cheatsheet](https://blog.anoff.io/puml-cheatsheet.pdf)
* [PlantUML Theme Documentation](https://plantuml.com/theme)
* [Stackoverflow question on embedding PlantUML in Markdown](https://stackoverflow.com/questions/32203610/how-to-integrate-uml-diagrams-into-gitlab-or-github)