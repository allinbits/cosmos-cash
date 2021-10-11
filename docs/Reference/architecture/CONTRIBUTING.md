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

The source code for all diagrams SHOULD be stored in `./src` folder AND have the `.puml` suffix. See the [PlanUML documentation](https://plantuml.com) for further details.

## Themes

Where possible all diagrams are rendered to have a similar look and feel. To achieve this look and feel, add the following theme to your diagram:

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

The [makefile](./makefile) builds all `src/*.puml` files. To create SVG images, run `make (svg|png)` as needed. Run 
`make clean` to get a clean folder ready for a new build.

Using PlantUML and makefile is the preferred and cleanest way to create all the architecture documentation.

### Proxy service

This method uses PlantUML's [proxy service](https://plantuml.com/server) to put generate images and embed into a Markdown page. 

To use the proxy service integration, use:

```
![cached image](http://www.plantuml.com/plantuml/proxy?src=https://raw.github.com/plantuml/plantuml-server/master/src/main/webapp/resource/test2diagrams.txt)
```
Or if caching is required:
```
![uncached image](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.github.com/plantuml/plantuml-server/master/src/main/webapp/resource/test2diagrams.txt)
```

### Visual Studio Code

If you use [Visual Studio Code](https://code.visualstudio.com/), a [PlantUML plugin](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml) is available. 

From here, you can preview diagrams by using the command palette (CTRL+SHIFT+P on Windows/Linux or COMMAND+SHIFT+P on macOS).

```
> PlantUML: Preview Current Diagram
```



## Reference

* [Markdown native diagrams with PlantUML](https://blog.anoff.io/2018-07-31-diagrams-with-plantuml/)
* [PlantUML Cheatsheet](https://blog.anoff.io/puml-cheatsheet.pdf)
* [PlantUML Theme Documentation](https://plantuml.com/theme)
* [Stackoverflow question on embedding PlantUML in Markdown](https://stackoverflow.com/questions/32203610/how-to-integrate-uml-diagrams-into-gitlab-or-github)